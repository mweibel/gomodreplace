package moduleone

import (
	"fmt"

	"github.com/mweibel/gomodreplace/lib"
)

func ModuleOne() {
	fmt.Println("Module One uses", lib.Library())
}
