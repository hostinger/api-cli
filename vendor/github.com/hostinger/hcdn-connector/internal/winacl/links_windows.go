//go:build windows

package winacl

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

// ErrRemote refuses a file or directory on a network volume.
var ErrRemote = errors.New("on a network volume, whose server reports any owner and DACL it likes; use a local disk")

// maxLinks is the number of reparse points that the I/O manager follows in one path.
const maxLinks = 63

// nameSurrogate marks a reparse tag that names another file or directory
// (IsReparseTagNameSurrogate). Other tags, as on a cloud placeholder, hold data.
const nameSurrogate = 0x20000000

// symlinkRelative: SYMLINK_FLAG_RELATIVE, the target is relative to the directory of the link.
const symlinkRelative = 1

// linkMask is writeMask and FILE_WRITE_ATTRIBUTES: FSCTL_SET_REPARSE_POINT
// re-points a link with either write access.
const linkMask = writeMask | windows.FILE_WRITE_ATTRIBUTES

// OpenTrustedPath opens path to judge it with CheckTrusted or CheckPrivateDir.
// os.Open follows every junction and symbolic link in a path, and any account
// can create a junction where it can create a directory (under %ProgramData%,
// in C:\) and re-point it at will, so the link would be judged by its target.
// OpenTrustedPath walks path from its volume root and opens each component
// without following reparse points. A reparse point is refused unless its owner
// is trusted and its DACL lets no one else change it, as in CheckTrusted; the
// target of a junction or symbolic link that passes is walked the same way.
// Components that are not reparse points need no trusted owner: C:\ and
// C:\ProgramData let any user add a directory. Every component must be on a
// local volume (ErrRemote).
//
// The file is named by the path walked, links resolved: its directory is the
// one that holds it.
func OpenTrustedPath(path string) (*os.File, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	vol := filepath.VolumeName(abs)
	cur := vol + `\`
	todo := strings.Split(abs[len(vol):], `\`)
	var via string // the last link followed, for the error
	// f is the handle of cur, or nil after a link or "..".
	f, _, _, err := walkComponent(cur)
	if err != nil {
		return nil, err
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	for links := 0; len(todo) > 0; {
		name := todo[0]
		todo = todo[1:]
		switch name {
		case "", ".":
			continue
		case "..":
			cur = filepath.Dir(cur) // cur holds no link: its parent is the real one
			if f != nil {
				f.Close()
				f = nil
			}
			continue
		}
		p := filepath.Join(cur, name)
		next, target, relative, err := walkComponent(p)
		if err != nil && via != "" {
			return nil, fmt.Errorf("through link %s: %w", via, err)
		}
		if err != nil {
			return nil, err
		}
		if f != nil {
			f.Close()
		}
		f = next
		if target == "" {
			cur = p
			continue
		}
		if links++; links > maxLinks {
			return nil, fmt.Errorf("%s: more than %d links", path, maxLinks)
		}
		via = p
		switch {
		case !relative:
			if vol = filepath.VolumeName(target); vol == "" {
				return nil, fmt.Errorf("link %s: target %s has no volume", via, target)
			}
			cur = vol + `\`
			target = target[len(vol):]
			if f, _, _, err = walkComponent(cur); err != nil {
				return nil, fmt.Errorf("through link %s: %w", via, err)
			}
		case strings.HasPrefix(target, `\`):
			cur = filepath.VolumeName(cur) + `\`
		}
		todo = append(strings.Split(target, `\`), todo...)
	}
	if f == nil {
		g, target, _, err := walkComponent(cur)
		if err != nil {
			return nil, err
		}
		if target != "" {
			return nil, fmt.Errorf("%s: changed while it was checked", cur)
		}
		f = g
	}
	g := f
	f = nil
	return g, nil
}

// walkComponent opens p without following it. It returns the open file, or
// the target of a junction or symbolic link that passes the trust check.
func walkComponent(p string) (f *os.File, target string, relative bool, err error) {
	f, err = openNoFollow(p)
	if err != nil {
		return nil, "", false, err
	}
	target, relative, err = linkTarget(f)
	if err != nil || target != "" {
		f.Close()
		return nil, target, relative, err
	}
	return f, "", false, nil
}

// linkTarget returns the target of f when it is a junction or a symbolic link.
// A reparse point of any tag must pass the trust check.
func linkTarget(f *os.File) (target string, relative bool, err error) {
	var fi windows.ByHandleFileInformation
	if err := windows.GetFileInformationByHandle(windows.Handle(f.Fd()), &fi); err != nil {
		return "", false, &os.PathError{Op: "stat", Path: f.Name(), Err: err}
	}
	if fi.FileAttributes&windows.FILE_ATTRIBUTE_REPARSE_POINT == 0 {
		return "", false, nil
	}
	if err := checkOwnerAndDACL(f, linkMask, "write access"); err != nil {
		return "", false, fmt.Errorf("link %s: %w", f.Name(), err)
	}
	buf := make([]byte, windows.MAXIMUM_REPARSE_DATA_BUFFER_SIZE)
	var n uint32
	if err := windows.DeviceIoControl(windows.Handle(f.Fd()), windows.FSCTL_GET_REPARSE_POINT, nil, 0, &buf[0], uint32(len(buf)), &n, nil); err != nil {
		return "", false, fmt.Errorf("link %s: %w", f.Name(), err)
	}
	tag, target, relative, err := parseReparse(buf[:n])
	switch {
	case err != nil:
		return "", false, fmt.Errorf("link %s: %w", f.Name(), err)
	case tag&nameSurrogate != 0 && target == "":
		return "", false, fmt.Errorf("link %s: reparse tag %#x is not a junction or a symbolic link", f.Name(), tag)
	}
	return target, relative, nil
}

// parseReparse reads a REPARSE_DATA_BUFFER. The target of a junction or a
// symbolic link is returned as a Win32 path, or as given when it is relative;
// other tags return no target.
func parseReparse(b []byte) (tag uint32, target string, relative bool, err error) {
	if len(b) < 8 {
		return 0, "", false, errors.New("short reparse data")
	}
	tag = binary.LittleEndian.Uint32(b)
	b = b[8:]
	var names int // offset of PathBuffer
	switch tag {
	case windows.IO_REPARSE_TAG_MOUNT_POINT:
		names = 8
	case windows.IO_REPARSE_TAG_SYMLINK:
		names = 12
	default:
		return tag, "", false, nil
	}
	if len(b) < names {
		return 0, "", false, errors.New("short reparse data")
	}
	off, n := int(binary.LittleEndian.Uint16(b)), int(binary.LittleEndian.Uint16(b[2:]))
	relative = tag == windows.IO_REPARSE_TAG_SYMLINK && binary.LittleEndian.Uint32(b[8:])&symlinkRelative != 0
	b = b[names:]
	if off%2 != 0 || n%2 != 0 || n == 0 || off+n > len(b) {
		return 0, "", false, errors.New("bad reparse data")
	}
	u := make([]uint16, n/2)
	for i := range u {
		u[i] = binary.LittleEndian.Uint16(b[off+2*i:])
	}
	target = windows.UTF16ToString(u)
	if relative {
		return tag, target, true, nil
	}
	if target, err = win32Path(target); err != nil {
		return 0, "", false, err
	}
	return tag, target, false, nil
}

// win32Path turns the NT path that a link stores (\??\C:\dir, \??\UNC\server\share,
// \??\Volume{guid}\ for a mount point) into a Win32 path.
func win32Path(nt string) (string, error) {
	rest, ok := strings.CutPrefix(nt, `\??\`)
	if !ok {
		return "", fmt.Errorf("target %s is not a Win32 path", nt)
	}
	if share, ok := strings.CutPrefix(rest, `UNC\`); ok {
		return `\\` + share, nil
	}
	if len(rest) >= 2 && rest[1] == ':' {
		return rest, nil
	}
	return `\\?\` + rest, nil
}

// openNoFollow opens p to read its attributes and security descriptor, and
// not the target when it is a reparse point.
func openNoFollow(p string) (*os.File, error) {
	u, err := windows.UTF16PtrFromString(p)
	if err != nil {
		return nil, err
	}
	h, err := windows.CreateFile(u, windows.READ_CONTROL|windows.FILE_READ_ATTRIBUTES,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE|windows.FILE_SHARE_DELETE, nil, windows.OPEN_EXISTING,
		windows.FILE_FLAG_OPEN_REPARSE_POINT|windows.FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return nil, &os.PathError{Op: "open", Path: p, Err: err}
	}
	f := os.NewFile(uintptr(h), p)
	if err := checkLocal(h); err != nil {
		f.Close()
		return nil, fmt.Errorf("%s: %w", p, err)
	}
	return f, nil
}

var procNtQueryVolumeInformationFile = windows.NewLazySystemDLL("ntdll.dll").NewProc("NtQueryVolumeInformationFile")

// FILE_FS_DEVICE_INFORMATION and the values that mark a network volume.
const (
	fileFsDeviceInformation     = 4
	fileRemoteDevice            = 0x10
	fileDeviceDfs               = 0x06
	fileDeviceNetwork           = 0x12
	fileDeviceNetworkFileSystem = 0x14
	fileDeviceNetworkRedirector = 0x28
	fileDeviceDfsFileSystem     = 0x35
	fileDeviceDfsVolume         = 0x36
)

// checkLocal refuses a handle on a network volume (an SMB or WebDAV share, a
// mapped drive, DFS). The device of the open handle decides, not the path: a
// drive letter or a symbolic link can stand for a share.
func checkLocal(h windows.Handle) error {
	var info struct{ DeviceType, Characteristics uint32 }
	var iosb windows.IO_STATUS_BLOCK
	r, _, _ := procNtQueryVolumeInformationFile.Call(uintptr(h), uintptr(unsafe.Pointer(&iosb)),
		uintptr(unsafe.Pointer(&info)), unsafe.Sizeof(info), fileFsDeviceInformation)
	if st := windows.NTStatus(r); st != windows.STATUS_SUCCESS {
		return fmt.Errorf("query the volume: %w", st)
	}
	if remoteDevice(info.DeviceType, info.Characteristics) {
		return ErrRemote
	}
	return nil
}

func remoteDevice(typ, characteristics uint32) bool {
	switch typ {
	case fileDeviceDfs, fileDeviceNetwork, fileDeviceNetworkFileSystem, fileDeviceNetworkRedirector,
		fileDeviceDfsFileSystem, fileDeviceDfsVolume:
		return true
	}
	return characteristics&fileRemoteDevice != 0
}
