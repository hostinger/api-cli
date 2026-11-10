//go:build windows

package winacl

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

// accessMask is any right that lets a trustee read or replace the file content,
// directly or after it rewrites the DACL or takes ownership (Unix: perm&0o077).
const accessMask = windows.FILE_READ_DATA | windows.GENERIC_READ | windows.GENERIC_ALL |
	windows.FILE_WRITE_DATA | windows.FILE_APPEND_DATA | windows.GENERIC_WRITE | windows.DELETE |
	windows.WRITE_DAC | windows.WRITE_OWNER

// ACE types that x/sys does not name. A callback ACE has the ACCESS_ALLOWED_ACE layout.
const (
	aceAllowedObject         = 0x5
	aceDeniedObject          = 0x6
	aceAllowedCallback       = 0x9
	aceDeniedCallback        = 0xA
	aceAllowedCallbackObject = 0xB
	aceDeniedCallbackObject  = 0xC
)

var broadGroups = []struct {
	sid  windows.WELL_KNOWN_SID_TYPE
	name string
}{
	{windows.WinWorldSid, "Everyone"},
	{windows.WinBuiltinUsersSid, "Users"},
	{windows.WinAuthenticatedUserSid, "Authenticated Users"},
	{windows.WinBuiltinGuestsSid, "Guests"},
	{windows.WinInteractiveSid, "INTERACTIVE"},
	{windows.WinNetworkSid, "NETWORK"},
	{windows.WinAnonymousSid, "ANONYMOUS LOGON"},
	{windows.WinLocalSid, "LOCAL"},
	{windows.WinBuiltinAnyPackageSid, "ALL APPLICATION PACKAGES"},
}

// Grantee is a broad group that a DACL lets read or write a file.
type Grantee struct{ Name, SID string }

// domainGroup names the domain-relative groups (S-1-5-21-<domain>-<rid>) that hold every account.
func domainGroup(sid string) string {
	if !strings.HasPrefix(sid, "S-1-5-21-") {
		return ""
	}
	switch {
	case strings.HasSuffix(sid, "-513"):
		return "Domain Users"
	case strings.HasSuffix(sid, "-514"):
		return "Domain Guests"
	}
	return ""
}

// PrivateSDDL grants full access to the current user and SYSTEM, without inherited ACEs.
func PrivateSDDL() (string, error) {
	u, err := windows.GetCurrentProcessToken().GetTokenUser()
	if err != nil {
		return "", err
	}
	return "D:P(A;;FA;;;" + u.User.Sid.String() + ")(A;;FA;;;SY)", nil
}

func attrs(sddl string) (*windows.SecurityAttributes, error) {
	sd, err := windows.SecurityDescriptorFromString(sddl)
	if err != nil {
		return nil, err
	}
	return &windows.SecurityAttributes{Length: uint32(unsafe.Sizeof(windows.SecurityAttributes{})), SecurityDescriptor: sd}, nil
}

// CreatePrivate creates path for writing with PrivateSDDL; it fails when path exists.
func CreatePrivate(path string) (*os.File, error) {
	sddl, err := PrivateSDDL()
	if err != nil {
		return nil, err
	}
	sa, err := attrs(sddl)
	if err != nil {
		return nil, err
	}
	p, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}
	h, err := windows.CreateFile(p, windows.GENERIC_WRITE, 0, sa, windows.CREATE_NEW, windows.FILE_ATTRIBUTE_NORMAL, 0)
	if err != nil {
		return nil, &os.PathError{Op: "open", Path: path, Err: err}
	}
	return os.NewFile(uintptr(h), path), nil
}

// MkdirProtected creates dir with sddl. An existing directory keeps its DACL.
func MkdirProtected(dir, sddl string) error {
	sa, err := attrs(sddl)
	if err != nil {
		return err
	}
	p, err := windows.UTF16PtrFromString(dir)
	if err != nil {
		return err
	}
	if err := windows.CreateDirectory(p, sa); err != nil && !errors.Is(err, windows.ERROR_ALREADY_EXISTS) {
		return &os.PathError{Op: "mkdir", Path: dir, Err: err}
	}
	return nil
}

// Set replaces the DACL of path with the one in sddl and drops inherited ACEs.
func Set(path, sddl string) error {
	sd, err := windows.SecurityDescriptorFromString(sddl)
	if err != nil {
		return err
	}
	dacl, _, err := sd.DACL()
	if err != nil {
		return err
	}
	return windows.SetNamedSecurityInfo(path, windows.SE_FILE_OBJECT,
		windows.DACL_SECURITY_INFORMATION|windows.PROTECTED_DACL_SECURITY_INFORMATION, nil, nil, dacl, nil)
}

// Restrict sets PrivateSDDL on path.
func Restrict(path string) error {
	sddl, err := PrivateSDDL()
	if err != nil {
		return err
	}
	return Set(path, sddl)
}

// BroadAccess lists the broad groups (Everyone, Users, ...) that the DACL of f lets read or write it.
func BroadAccess(f *os.File) ([]Grantee, error) {
	sd, err := windows.GetSecurityInfo(windows.Handle(f.Fd()), windows.SE_FILE_OBJECT, windows.DACL_SECURITY_INFORMATION)
	if err != nil {
		return nil, err
	}
	dacl, _, err := sd.DACL()
	if errors.Is(err, windows.ERROR_OBJECT_NOT_FOUND) || (err == nil && dacl == nil) {
		return []Grantee{{Name: "Everyone (no DACL)", SID: "S-1-1-0"}}, nil // a NULL DACL grants everything
	}
	if err != nil {
		return nil, err
	}
	sids := make([]*windows.SID, len(broadGroups))
	for i, g := range broadGroups {
		if sids[i], err = windows.CreateWellKnownSid(g.sid); err != nil {
			return nil, err
		}
	}
	var out []Grantee
	for i := uint32(0); i < uint32(dacl.AceCount); i++ {
		var ace *windows.ACCESS_ALLOWED_ACE
		if err := windows.GetAce(dacl, i, &ace); err != nil {
			return nil, err
		}
		t := ace.Header.AceType
		if (t != windows.ACCESS_ALLOWED_ACE_TYPE && t != aceAllowedCallback) || ace.Header.AceFlags&windows.INHERIT_ONLY_ACE != 0 ||
			uint32(ace.Mask)&accessMask == 0 {
			continue
		}
		sid := (*windows.SID)(unsafe.Pointer(&ace.SidStart))
		s := sid.String()
		if name := domainGroup(s); name != "" {
			out = append(out, Grantee{Name: name, SID: s})
			continue
		}
		for j, g := range broadGroups {
			if sid.Equals(sids[j]) {
				out = append(out, Grantee{Name: g.name, SID: s})
			}
		}
	}
	return out, nil
}

// writeMask is any right that lets a trustee change a file or the entries of a
// directory: write or append (on a directory: add a file or a subdirectory),
// delete it or a child, or rewrite its DACL or owner.
const writeMask = windows.FILE_WRITE_DATA | windows.FILE_APPEND_DATA | windows.GENERIC_WRITE | windows.GENERIC_ALL |
	windows.DELETE | fileDeleteChild | windows.WRITE_DAC | windows.WRITE_OWNER

// fileDeleteChild: x/sys does not name FILE_DELETE_CHILD.
const fileDeleteChild = 0x40

// CheckPrivateDir refuses a directory that another account owns, or whose DACL
// grants access to anyone but SYSTEM, Administrators, TrustedInstaller and the
// current user. Inherit-only ACEs count: the files created in the directory get
// them. A directory on a network volume is refused (ErrRemote).
func CheckPrivateDir(f *os.File) error { return checkOwnerAndDACL(f, ^uint32(0), "access") }

// CheckTrusted refuses a file or directory that another account owns, or whose
// DACL lets anyone but SYSTEM, Administrators, TrustedInstaller and the current
// user change it (writeMask). Read access is allowed. Inherit-only ACEs count:
// the files created in a directory get them. A file on a network volume is
// refused (ErrRemote). f is judged as opened: see OpenTrustedPath for links.
func CheckTrusted(f *os.File) error { return checkOwnerAndDACL(f, writeMask, "write access") }

func checkOwnerAndDACL(f *os.File, mask uint32, what string) error {
	if err := checkLocal(windows.Handle(f.Fd())); err != nil {
		return err
	}
	sd, err := windows.GetSecurityInfo(windows.Handle(f.Fd()), windows.SE_FILE_OBJECT,
		windows.OWNER_SECURITY_INFORMATION|windows.DACL_SECURITY_INFORMATION)
	if err != nil {
		return err
	}
	trusted, err := trustedSIDs()
	if err != nil {
		return err
	}
	owner, _, err := sd.Owner()
	if err != nil {
		return err
	}
	if !containsSID(trusted, owner) {
		return fmt.Errorf("owned by %s", accountName(owner))
	}
	dacl, _, err := sd.DACL()
	if errors.Is(err, windows.ERROR_OBJECT_NOT_FOUND) || (err == nil && dacl == nil) {
		return errors.New("no DACL: every account has full access")
	}
	if err != nil {
		return err
	}
	for i := uint32(0); i < uint32(dacl.AceCount); i++ {
		var ace *windows.ACCESS_ALLOWED_ACE
		if err := windows.GetAce(dacl, i, &ace); err != nil {
			return err
		}
		switch ace.Header.AceType {
		case windows.ACCESS_DENIED_ACE_TYPE, aceDeniedObject, aceDeniedCallback, aceDeniedCallbackObject:
			continue
		case windows.ACCESS_ALLOWED_ACE_TYPE, aceAllowedCallback:
			if sid := (*windows.SID)(unsafe.Pointer(&ace.SidStart)); uint32(ace.Mask)&mask != 0 && !containsSID(trusted, sid) {
				return fmt.Errorf("grants %s to %s", what, accountName(sid))
			}
		case aceAllowedObject, aceAllowedCallbackObject:
			if uint32(ace.Mask)&mask != 0 {
				return errors.New("has an object ACE")
			}
		default:
			return fmt.Errorf("has an ACE of type %d", ace.Header.AceType)
		}
	}
	return nil
}

// trustedInstallerSID is NT SERVICE\TrustedInstaller, which holds Full Control
// on C:\Program Files and passes it to every directory created there. It
// outranks Administrators, so trusting it weakens nothing.
const trustedInstallerSID = "S-1-5-80-956008885-3418522649-1831038044-1853292631-2271478464"

// trustedSIDs: SYSTEM, Administrators, TrustedInstaller, the current user, and
// CREATOR OWNER and OWNER RIGHTS, which resolve to the owner that
// checkOwnerAndDACL already checked.
func trustedSIDs() ([]*windows.SID, error) {
	u, err := windows.GetCurrentProcessToken().GetTokenUser()
	if err != nil {
		return nil, err
	}
	ti, err := windows.StringToSid(trustedInstallerSID)
	if err != nil {
		return nil, err
	}
	out := []*windows.SID{u.User.Sid, ti}
	for _, t := range []windows.WELL_KNOWN_SID_TYPE{windows.WinLocalSystemSid, windows.WinBuiltinAdministratorsSid,
		windows.WinCreatorOwnerSid, windows.WinCreatorOwnerRightsSid} {
		s, err := windows.CreateWellKnownSid(t)
		if err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, nil
}

func containsSID(list []*windows.SID, sid *windows.SID) bool {
	for _, s := range list {
		if sid.Equals(s) {
			return true
		}
	}
	return false
}

func accountName(sid *windows.SID) string {
	if account, domain, _, err := sid.LookupAccount(""); err == nil {
		if domain != "" {
			return domain + `\` + account
		}
		return account
	}
	return sid.String()
}

// HasOnly reports whether every allow ACE of path names one of sids (tests).
func HasOnly(path string, sids ...*windows.SID) (bool, error) {
	sd, err := windows.GetNamedSecurityInfo(path, windows.SE_FILE_OBJECT, windows.DACL_SECURITY_INFORMATION)
	if err != nil {
		return false, err
	}
	dacl, _, err := sd.DACL()
	if err != nil || dacl == nil {
		return false, err
	}
	for i := uint32(0); i < uint32(dacl.AceCount); i++ {
		var ace *windows.ACCESS_ALLOWED_ACE
		if err := windows.GetAce(dacl, i, &ace); err != nil {
			return false, err
		}
		if ace.Header.AceType != windows.ACCESS_ALLOWED_ACE_TYPE {
			continue
		}
		sid := (*windows.SID)(unsafe.Pointer(&ace.SidStart))
		ok := false
		for _, s := range sids {
			if sid.Equals(s) {
				ok = true
			}
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}
