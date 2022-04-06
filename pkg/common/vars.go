package common

const CHKSIZE uint32 = 1 << 22
const MAX_IO_SIZE uint32 = 4096

const IOV_MAX uint32 = CHKSIZE / MAX_IO_SIZE

const MaxMsgSize uint32 = CHKSIZE + 4096
