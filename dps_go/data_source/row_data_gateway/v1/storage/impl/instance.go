package impl

import (
	"rdg_v1/storage"
	"rdg_v1/storage/impl/memory"
)

// singleton?
func Instance(value storage.Type) storage.Storage {
	if value == storage.MEMORY {
		return memory.Instance()
	} else if value == storage.SQL_POSTGRES {
		//...
	} else if value == storage.SQL_LITE {
		//...
	}

	return memory.Instance()
}
