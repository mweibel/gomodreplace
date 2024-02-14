package moduletwo

import (
	"fmt"

	"github.com/mweibel/gomodreplace/lib"
)

func ModuleTwo() {
	fmt.Println("Module Two uses", lib.Library())
}
