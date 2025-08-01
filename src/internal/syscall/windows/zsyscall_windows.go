// Code generated by 'go generate'; DO NOT EDIT.

package windows

import (
	"internal/syscall/windows/sysdll"
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modadvapi32         = syscall.NewLazyDLL(sysdll.Add("advapi32.dll"))
	modbcryptprimitives = syscall.NewLazyDLL(sysdll.Add("bcryptprimitives.dll"))
	modiphlpapi         = syscall.NewLazyDLL(sysdll.Add("iphlpapi.dll"))
	modkernel32         = syscall.NewLazyDLL(sysdll.Add("kernel32.dll"))
	modnetapi32         = syscall.NewLazyDLL(sysdll.Add("netapi32.dll"))
	modntdll            = syscall.NewLazyDLL(sysdll.Add("ntdll.dll"))
	modpsapi            = syscall.NewLazyDLL(sysdll.Add("psapi.dll"))
	moduserenv          = syscall.NewLazyDLL(sysdll.Add("userenv.dll"))
	modws2_32           = syscall.NewLazyDLL(sysdll.Add("ws2_32.dll"))

	procAdjustTokenPrivileges             = modadvapi32.NewProc("AdjustTokenPrivileges")
	procDuplicateTokenEx                  = modadvapi32.NewProc("DuplicateTokenEx")
	procGetSidIdentifierAuthority         = modadvapi32.NewProc("GetSidIdentifierAuthority")
	procGetSidSubAuthority                = modadvapi32.NewProc("GetSidSubAuthority")
	procGetSidSubAuthorityCount           = modadvapi32.NewProc("GetSidSubAuthorityCount")
	procImpersonateLoggedOnUser           = modadvapi32.NewProc("ImpersonateLoggedOnUser")
	procImpersonateSelf                   = modadvapi32.NewProc("ImpersonateSelf")
	procInitializeAcl                     = modadvapi32.NewProc("InitializeAcl")
	procIsValidSid                        = modadvapi32.NewProc("IsValidSid")
	procLogonUserW                        = modadvapi32.NewProc("LogonUserW")
	procLookupPrivilegeValueW             = modadvapi32.NewProc("LookupPrivilegeValueW")
	procOpenSCManagerW                    = modadvapi32.NewProc("OpenSCManagerW")
	procOpenServiceW                      = modadvapi32.NewProc("OpenServiceW")
	procOpenThreadToken                   = modadvapi32.NewProc("OpenThreadToken")
	procQueryServiceStatus                = modadvapi32.NewProc("QueryServiceStatus")
	procRevertToSelf                      = modadvapi32.NewProc("RevertToSelf")
	procSetNamedSecurityInfoW             = modadvapi32.NewProc("SetNamedSecurityInfoW")
	procSetTokenInformation               = modadvapi32.NewProc("SetTokenInformation")
	procProcessPrng                       = modbcryptprimitives.NewProc("ProcessPrng")
	procGetAdaptersAddresses              = modiphlpapi.NewProc("GetAdaptersAddresses")
	procCreateEventW                      = modkernel32.NewProc("CreateEventW")
	procCreateIoCompletionPort            = modkernel32.NewProc("CreateIoCompletionPort")
	procCreateNamedPipeW                  = modkernel32.NewProc("CreateNamedPipeW")
	procGetACP                            = modkernel32.NewProc("GetACP")
	procGetComputerNameExW                = modkernel32.NewProc("GetComputerNameExW")
	procGetConsoleCP                      = modkernel32.NewProc("GetConsoleCP")
	procGetCurrentThread                  = modkernel32.NewProc("GetCurrentThread")
	procGetFileInformationByHandleEx      = modkernel32.NewProc("GetFileInformationByHandleEx")
	procGetFinalPathNameByHandleW         = modkernel32.NewProc("GetFinalPathNameByHandleW")
	procGetModuleFileNameW                = modkernel32.NewProc("GetModuleFileNameW")
	procGetModuleHandleW                  = modkernel32.NewProc("GetModuleHandleW")
	procGetOverlappedResult               = modkernel32.NewProc("GetOverlappedResult")
	procGetTempPath2W                     = modkernel32.NewProc("GetTempPath2W")
	procGetVolumeInformationByHandleW     = modkernel32.NewProc("GetVolumeInformationByHandleW")
	procGetVolumeNameForVolumeMountPointW = modkernel32.NewProc("GetVolumeNameForVolumeMountPointW")
	procLockFileEx                        = modkernel32.NewProc("LockFileEx")
	procModule32FirstW                    = modkernel32.NewProc("Module32FirstW")
	procModule32NextW                     = modkernel32.NewProc("Module32NextW")
	procMoveFileExW                       = modkernel32.NewProc("MoveFileExW")
	procMultiByteToWideChar               = modkernel32.NewProc("MultiByteToWideChar")
	procRtlLookupFunctionEntry            = modkernel32.NewProc("RtlLookupFunctionEntry")
	procRtlVirtualUnwind                  = modkernel32.NewProc("RtlVirtualUnwind")
	procSetFileInformationByHandle        = modkernel32.NewProc("SetFileInformationByHandle")
	procUnlockFileEx                      = modkernel32.NewProc("UnlockFileEx")
	procVirtualQuery                      = modkernel32.NewProc("VirtualQuery")
	procNetShareAdd                       = modnetapi32.NewProc("NetShareAdd")
	procNetShareDel                       = modnetapi32.NewProc("NetShareDel")
	procNetUserAdd                        = modnetapi32.NewProc("NetUserAdd")
	procNetUserDel                        = modnetapi32.NewProc("NetUserDel")
	procNetUserGetLocalGroups             = modnetapi32.NewProc("NetUserGetLocalGroups")
	procNtCreateFile                      = modntdll.NewProc("NtCreateFile")
	procNtOpenFile                        = modntdll.NewProc("NtOpenFile")
	procNtQueryInformationFile            = modntdll.NewProc("NtQueryInformationFile")
	procNtSetInformationFile              = modntdll.NewProc("NtSetInformationFile")
	procRtlGetVersion                     = modntdll.NewProc("RtlGetVersion")
	procRtlIsDosDeviceName_U              = modntdll.NewProc("RtlIsDosDeviceName_U")
	procRtlNtStatusToDosErrorNoTeb        = modntdll.NewProc("RtlNtStatusToDosErrorNoTeb")
	procGetProcessMemoryInfo              = modpsapi.NewProc("GetProcessMemoryInfo")
	procCreateEnvironmentBlock            = moduserenv.NewProc("CreateEnvironmentBlock")
	procDestroyEnvironmentBlock           = moduserenv.NewProc("DestroyEnvironmentBlock")
	procGetProfilesDirectoryW             = moduserenv.NewProc("GetProfilesDirectoryW")
	procWSADuplicateSocketW               = modws2_32.NewProc("WSADuplicateSocketW")
	procWSAGetOverlappedResult            = modws2_32.NewProc("WSAGetOverlappedResult")
	procWSASocketW                        = modws2_32.NewProc("WSASocketW")
)

func adjustTokenPrivileges(token syscall.Token, disableAllPrivileges bool, newstate *TOKEN_PRIVILEGES, buflen uint32, prevstate *TOKEN_PRIVILEGES, returnlen *uint32) (ret uint32, err error) {
	var _p0 uint32
	if disableAllPrivileges {
		_p0 = 1
	}
	r0, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(), 6, uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(newstate)), uintptr(buflen), uintptr(unsafe.Pointer(prevstate)), uintptr(unsafe.Pointer(returnlen)))
	ret = uint32(r0)
	if true {
		err = errnoErr(e1)
	}
	return
}

func DuplicateTokenEx(hExistingToken syscall.Token, dwDesiredAccess uint32, lpTokenAttributes *syscall.SecurityAttributes, impersonationLevel uint32, tokenType TokenType, phNewToken *syscall.Token) (err error) {
	r1, _, e1 := syscall.Syscall6(procDuplicateTokenEx.Addr(), 6, uintptr(hExistingToken), uintptr(dwDesiredAccess), uintptr(unsafe.Pointer(lpTokenAttributes)), uintptr(impersonationLevel), uintptr(tokenType), uintptr(unsafe.Pointer(phNewToken)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getSidIdentifierAuthority(sid *syscall.SID) (idauth uintptr) {
	r0, _, _ := syscall.Syscall(procGetSidIdentifierAuthority.Addr(), 1, uintptr(unsafe.Pointer(sid)), 0, 0)
	idauth = uintptr(r0)
	return
}

func getSidSubAuthority(sid *syscall.SID, subAuthorityIdx uint32) (subAuth uintptr) {
	r0, _, _ := syscall.Syscall(procGetSidSubAuthority.Addr(), 2, uintptr(unsafe.Pointer(sid)), uintptr(subAuthorityIdx), 0)
	subAuth = uintptr(r0)
	return
}

func getSidSubAuthorityCount(sid *syscall.SID) (count uintptr) {
	r0, _, _ := syscall.Syscall(procGetSidSubAuthorityCount.Addr(), 1, uintptr(unsafe.Pointer(sid)), 0, 0)
	count = uintptr(r0)
	return
}

func ImpersonateLoggedOnUser(token syscall.Token) (err error) {
	r1, _, e1 := syscall.Syscall(procImpersonateLoggedOnUser.Addr(), 1, uintptr(token), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ImpersonateSelf(impersonationlevel uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procImpersonateSelf.Addr(), 1, uintptr(impersonationlevel), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func InitializeAcl(acl *ACL, length uint32, revision uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procInitializeAcl.Addr(), 3, uintptr(unsafe.Pointer(acl)), uintptr(length), uintptr(revision))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func IsValidSid(sid *syscall.SID) (valid bool) {
	r0, _, _ := syscall.Syscall(procIsValidSid.Addr(), 1, uintptr(unsafe.Pointer(sid)), 0, 0)
	valid = r0 != 0
	return
}

func LogonUser(username *uint16, domain *uint16, password *uint16, logonType uint32, logonProvider uint32, token *syscall.Token) (err error) {
	r1, _, e1 := syscall.Syscall6(procLogonUserW.Addr(), 6, uintptr(unsafe.Pointer(username)), uintptr(unsafe.Pointer(domain)), uintptr(unsafe.Pointer(password)), uintptr(logonType), uintptr(logonProvider), uintptr(unsafe.Pointer(token)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) (err error) {
	r1, _, e1 := syscall.Syscall(procLookupPrivilegeValueW.Addr(), 3, uintptr(unsafe.Pointer(systemname)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(luid)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procOpenSCManagerW.Addr(), 3, uintptr(unsafe.Pointer(machineName)), uintptr(unsafe.Pointer(databaseName)), uintptr(access))
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func OpenService(mgr syscall.Handle, serviceName *uint16, access uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procOpenServiceW.Addr(), 3, uintptr(mgr), uintptr(unsafe.Pointer(serviceName)), uintptr(access))
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func OpenThreadToken(h syscall.Handle, access uint32, openasself bool, token *syscall.Token) (err error) {
	var _p0 uint32
	if openasself {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procOpenThreadToken.Addr(), 4, uintptr(h), uintptr(access), uintptr(_p0), uintptr(unsafe.Pointer(token)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func QueryServiceStatus(hService syscall.Handle, lpServiceStatus *SERVICE_STATUS) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryServiceStatus.Addr(), 2, uintptr(hService), uintptr(unsafe.Pointer(lpServiceStatus)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func RevertToSelf() (err error) {
	r1, _, e1 := syscall.Syscall(procRevertToSelf.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetNamedSecurityInfo(objectName string, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *syscall.SID, group *syscall.SID, dacl *ACL, sacl *ACL) (ret error) {
	var _p0 *uint16
	_p0, ret = syscall.UTF16PtrFromString(objectName)
	if ret != nil {
		return
	}
	return _SetNamedSecurityInfo(_p0, objectType, securityInformation, owner, group, dacl, sacl)
}

func _SetNamedSecurityInfo(objectName *uint16, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *syscall.SID, group *syscall.SID, dacl *ACL, sacl *ACL) (ret error) {
	r0, _, _ := syscall.Syscall9(procSetNamedSecurityInfoW.Addr(), 7, uintptr(unsafe.Pointer(objectName)), uintptr(objectType), uintptr(securityInformation), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(sacl)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func SetTokenInformation(tokenHandle syscall.Token, tokenInformationClass uint32, tokenInformation unsafe.Pointer, tokenInformationLength uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetTokenInformation.Addr(), 4, uintptr(tokenHandle), uintptr(tokenInformationClass), uintptr(tokenInformation), uintptr(tokenInformationLength), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ProcessPrng(buf []byte) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r1, _, e1 := syscall.Syscall(procProcessPrng.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetAdaptersAddresses(family uint32, flags uint32, reserved unsafe.Pointer, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) {
	r0, _, _ := syscall.Syscall6(procGetAdaptersAddresses.Addr(), 5, uintptr(family), uintptr(flags), uintptr(reserved), uintptr(unsafe.Pointer(adapterAddresses)), uintptr(unsafe.Pointer(sizePointer)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func CreateEvent(eventAttrs *SecurityAttributes, manualReset uint32, initialState uint32, name *uint16) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateEventW.Addr(), 4, uintptr(unsafe.Pointer(eventAttrs)), uintptr(manualReset), uintptr(initialState), uintptr(unsafe.Pointer(name)), 0, 0)
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateIoCompletionPort(filehandle syscall.Handle, cphandle syscall.Handle, key uintptr, threadcnt uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateIoCompletionPort.Addr(), 4, uintptr(filehandle), uintptr(cphandle), uintptr(key), uintptr(threadcnt), 0, 0)
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateNamedPipe(name *uint16, flags uint32, pipeMode uint32, maxInstances uint32, outSize uint32, inSize uint32, defaultTimeout uint32, sa *syscall.SecurityAttributes) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateNamedPipeW.Addr(), 8, uintptr(unsafe.Pointer(name)), uintptr(flags), uintptr(pipeMode), uintptr(maxInstances), uintptr(outSize), uintptr(inSize), uintptr(defaultTimeout), uintptr(unsafe.Pointer(sa)), 0)
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func GetACP() (acp uint32) {
	r0, _, _ := syscall.Syscall(procGetACP.Addr(), 0, 0, 0, 0)
	acp = uint32(r0)
	return
}

func GetComputerNameEx(nameformat uint32, buf *uint16, n *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetComputerNameExW.Addr(), 3, uintptr(nameformat), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(n)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetConsoleCP() (ccp uint32) {
	r0, _, _ := syscall.Syscall(procGetConsoleCP.Addr(), 0, 0, 0, 0)
	ccp = uint32(r0)
	return
}

func GetCurrentThread() (pseudoHandle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetCurrentThread.Addr(), 0, 0, 0, 0)
	pseudoHandle = syscall.Handle(r0)
	if pseudoHandle == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetFileInformationByHandleEx(handle syscall.Handle, class uint32, info *byte, bufsize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetFileInformationByHandleEx.Addr(), 4, uintptr(handle), uintptr(class), uintptr(unsafe.Pointer(info)), uintptr(bufsize), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetFinalPathNameByHandle(file syscall.Handle, filePath *uint16, filePathSize uint32, flags uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetFinalPathNameByHandleW.Addr(), 4, uintptr(file), uintptr(unsafe.Pointer(filePath)), uintptr(filePathSize), uintptr(flags), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetModuleFileName(module syscall.Handle, fn *uint16, len uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleFileNameW.Addr(), 3, uintptr(module), uintptr(unsafe.Pointer(fn)), uintptr(len))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetModuleHandle(modulename *uint16) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleHandleW.Addr(), 1, uintptr(unsafe.Pointer(modulename)), 0, 0)
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetOverlappedResult(handle syscall.Handle, overlapped *syscall.Overlapped, done *uint32, wait bool) (err error) {
	var _p0 uint32
	if wait {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procGetOverlappedResult.Addr(), 4, uintptr(handle), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(done)), uintptr(_p0), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetTempPath2(buflen uint32, buf *uint16) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetTempPath2W.Addr(), 2, uintptr(buflen), uintptr(unsafe.Pointer(buf)), 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetVolumeInformationByHandle(file syscall.Handle, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetVolumeInformationByHandleW.Addr(), 8, uintptr(file), uintptr(unsafe.Pointer(volumeNameBuffer)), uintptr(volumeNameSize), uintptr(unsafe.Pointer(volumeNameSerialNumber)), uintptr(unsafe.Pointer(maximumComponentLength)), uintptr(unsafe.Pointer(fileSystemFlags)), uintptr(unsafe.Pointer(fileSystemNameBuffer)), uintptr(fileSystemNameSize), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetVolumeNameForVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16, bufferlength uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetVolumeNameForVolumeMountPointW.Addr(), 3, uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(unsafe.Pointer(volumeName)), uintptr(bufferlength))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func LockFileEx(file syscall.Handle, flags uint32, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procLockFileEx.Addr(), 6, uintptr(file), uintptr(flags), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func Module32First(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32FirstW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(moduleEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func Module32Next(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32NextW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(moduleEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func MoveFileEx(from *uint16, to *uint16, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procMoveFileExW.Addr(), 3, uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(to)), uintptr(flags))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func MultiByteToWideChar(codePage uint32, dwFlags uint32, str *byte, nstr int32, wchar *uint16, nwchar int32) (nwrite int32, err error) {
	r0, _, e1 := syscall.Syscall6(procMultiByteToWideChar.Addr(), 6, uintptr(codePage), uintptr(dwFlags), uintptr(unsafe.Pointer(str)), uintptr(nstr), uintptr(unsafe.Pointer(wchar)), uintptr(nwchar))
	nwrite = int32(r0)
	if nwrite == 0 {
		err = errnoErr(e1)
	}
	return
}

func RtlLookupFunctionEntry(pc uintptr, baseAddress *uintptr, table unsafe.Pointer) (ret *RUNTIME_FUNCTION) {
	r0, _, _ := syscall.Syscall(procRtlLookupFunctionEntry.Addr(), 3, uintptr(pc), uintptr(unsafe.Pointer(baseAddress)), uintptr(table))
	ret = (*RUNTIME_FUNCTION)(unsafe.Pointer(r0))
	return
}

func RtlVirtualUnwind(handlerType uint32, baseAddress uintptr, pc uintptr, entry *RUNTIME_FUNCTION, ctxt unsafe.Pointer, data unsafe.Pointer, frame *uintptr, ctxptrs unsafe.Pointer) (ret uintptr) {
	r0, _, _ := syscall.Syscall9(procRtlVirtualUnwind.Addr(), 8, uintptr(handlerType), uintptr(baseAddress), uintptr(pc), uintptr(unsafe.Pointer(entry)), uintptr(ctxt), uintptr(data), uintptr(unsafe.Pointer(frame)), uintptr(ctxptrs), 0)
	ret = uintptr(r0)
	return
}

func SetFileInformationByHandle(handle syscall.Handle, fileInformationClass uint32, buf unsafe.Pointer, bufsize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetFileInformationByHandle.Addr(), 4, uintptr(handle), uintptr(fileInformationClass), uintptr(buf), uintptr(bufsize), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func UnlockFileEx(file syscall.Handle, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procUnlockFileEx.Addr(), 5, uintptr(file), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualQuery.Addr(), 3, uintptr(address), uintptr(unsafe.Pointer(buffer)), uintptr(length))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func NetShareAdd(serverName *uint16, level uint32, buf *byte, parmErr *uint16) (neterr error) {
	r0, _, _ := syscall.Syscall6(procNetShareAdd.Addr(), 4, uintptr(unsafe.Pointer(serverName)), uintptr(level), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(parmErr)), 0, 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NetShareDel(serverName *uint16, netName *uint16, reserved uint32) (neterr error) {
	r0, _, _ := syscall.Syscall(procNetShareDel.Addr(), 3, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(netName)), uintptr(reserved))
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NetUserAdd(serverName *uint16, level uint32, buf *byte, parmErr *uint32) (neterr error) {
	r0, _, _ := syscall.Syscall6(procNetUserAdd.Addr(), 4, uintptr(unsafe.Pointer(serverName)), uintptr(level), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(parmErr)), 0, 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NetUserDel(serverName *uint16, userName *uint16) (neterr error) {
	r0, _, _ := syscall.Syscall(procNetUserDel.Addr(), 2, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(userName)), 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NetUserGetLocalGroups(serverName *uint16, userName *uint16, level uint32, flags uint32, buf **byte, prefMaxLen uint32, entriesRead *uint32, totalEntries *uint32) (neterr error) {
	r0, _, _ := syscall.Syscall9(procNetUserGetLocalGroups.Addr(), 8, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(userName)), uintptr(level), uintptr(flags), uintptr(unsafe.Pointer(buf)), uintptr(prefMaxLen), uintptr(unsafe.Pointer(entriesRead)), uintptr(unsafe.Pointer(totalEntries)), 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NtCreateFile(handle *syscall.Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, allocationSize *int64, attributes uint32, share uint32, disposition uint32, options uint32, eabuffer unsafe.Pointer, ealength uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall12(procNtCreateFile.Addr(), 11, uintptr(unsafe.Pointer(handle)), uintptr(access), uintptr(unsafe.Pointer(oa)), uintptr(unsafe.Pointer(iosb)), uintptr(unsafe.Pointer(allocationSize)), uintptr(attributes), uintptr(share), uintptr(disposition), uintptr(options), uintptr(eabuffer), uintptr(ealength), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func NtOpenFile(handle *syscall.Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, share uint32, options uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtOpenFile.Addr(), 6, uintptr(unsafe.Pointer(handle)), uintptr(access), uintptr(unsafe.Pointer(oa)), uintptr(unsafe.Pointer(iosb)), uintptr(share), uintptr(options))
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func NtQueryInformationFile(handle syscall.Handle, iosb *IO_STATUS_BLOCK, inBuffer unsafe.Pointer, inBufferLen uint32, class uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtQueryInformationFile.Addr(), 5, uintptr(handle), uintptr(unsafe.Pointer(iosb)), uintptr(inBuffer), uintptr(inBufferLen), uintptr(class), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func NtSetInformationFile(handle syscall.Handle, iosb *IO_STATUS_BLOCK, inBuffer unsafe.Pointer, inBufferLen uint32, class uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtSetInformationFile.Addr(), 5, uintptr(handle), uintptr(unsafe.Pointer(iosb)), uintptr(inBuffer), uintptr(inBufferLen), uintptr(class), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func rtlGetVersion(info *_OSVERSIONINFOEXW) {
	syscall.Syscall(procRtlGetVersion.Addr(), 1, uintptr(unsafe.Pointer(info)), 0, 0)
	return
}

func RtlIsDosDeviceName_U(name *uint16) (ret uint32) {
	r0, _, _ := syscall.Syscall(procRtlIsDosDeviceName_U.Addr(), 1, uintptr(unsafe.Pointer(name)), 0, 0)
	ret = uint32(r0)
	return
}

func rtlNtStatusToDosErrorNoTeb(ntstatus NTStatus) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procRtlNtStatusToDosErrorNoTeb.Addr(), 1, uintptr(ntstatus), 0, 0)
	ret = syscall.Errno(r0)
	return
}

func GetProcessMemoryInfo(handle syscall.Handle, memCounters *PROCESS_MEMORY_COUNTERS, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProcessMemoryInfo.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(memCounters)), uintptr(cb))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateEnvironmentBlock(block **uint16, token syscall.Token, inheritExisting bool) (err error) {
	var _p0 uint32
	if inheritExisting {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procCreateEnvironmentBlock.Addr(), 3, uintptr(unsafe.Pointer(block)), uintptr(token), uintptr(_p0))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func DestroyEnvironmentBlock(block *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyEnvironmentBlock.Addr(), 1, uintptr(unsafe.Pointer(block)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetProfilesDirectory(dir *uint16, dirLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProfilesDirectoryW.Addr(), 2, uintptr(unsafe.Pointer(dir)), uintptr(unsafe.Pointer(dirLen)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func WSADuplicateSocket(s syscall.Handle, processID uint32, info *syscall.WSAProtocolInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procWSADuplicateSocketW.Addr(), 3, uintptr(s), uintptr(processID), uintptr(unsafe.Pointer(info)))
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func WSAGetOverlappedResult(h syscall.Handle, o *syscall.Overlapped, bytes *uint32, wait bool, flags *uint32) (err error) {
	var _p0 uint32
	if wait {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procWSAGetOverlappedResult.Addr(), 5, uintptr(h), uintptr(unsafe.Pointer(o)), uintptr(unsafe.Pointer(bytes)), uintptr(_p0), uintptr(unsafe.Pointer(flags)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func WSASocket(af int32, typ int32, protocol int32, protinfo *syscall.WSAProtocolInfo, group uint32, flags uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procWSASocketW.Addr(), 6, uintptr(af), uintptr(typ), uintptr(protocol), uintptr(unsafe.Pointer(protinfo)), uintptr(group), uintptr(flags))
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		err = errnoErr(e1)
	}
	return
}
