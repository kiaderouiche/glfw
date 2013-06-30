package glfw3

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
import "C"

import (
	"errors"
	"unsafe"
)

//MakeContextCurrent makes the context of the window current.
func (w *Window) MakeContextCurrent() {
	C.glfwMakeContextCurrent(w.data)
}

//DetachCurrentContext detaches the current context.
func DetachCurrentContext() {
    C.glfwMakeContextCurrent(nil)
}

//GetCurrentContext returns the window whose context is current.
func GetCurrentContext() (*Window, error) {
	w := C.glfwGetCurrentContext()

	if w == nil {
		return nil, errors.New("Current context is not set.")
	}
	return &Window{w}, nil
}

//SwapBuffers swaps the front and back buffers of the window. If the
//swap interval is greater than zero, the GPU driver waits the specified number
//of screen updates before swapping the buffers.
func (w *Window) SwapBuffers() {
	C.glfwSwapBuffers(w.data)
}

//SwapInterval sets the swap interval for the current context, i.e. the number
//of screen updates to wait before swapping the buffers of a window and
//returning from SwapBuffers. This is sometimes called
//'vertical synchronization', 'vertical retrace synchronization' or 'vsync'.
//
//Contexts that support either of the WGL_EXT_swap_control_tear and
//GLX_EXT_swap_control_tear extensions also accept negative swap intervals,
//which allow the driver to swap even if a frame arrives a little bit late.
//You can check for the presence of these extensions using
//ExtensionSupported. For more information about swap tearing,
//see the extension specifications.
//
//Some GPU drivers do not honor the requested swap interval, either because of
//user settings that override the request or due to bugs in the driver.
func SwapInterval(interval int) {
	C.glfwSwapInterval(C.int(interval))
}

//ExtensionSupported returns whether the specified OpenGL or context creation
//API extension is supported by the current context. For example, on Windows
//both the OpenGL and WGL extension strings are checked.
//
//As this functions searches one or more extension strings on each call, it is
//recommended that you cache its results if it's going to be used frequently.
//The extension strings will not change during the lifetime of a context, so
//there is no danger in doing this.
func ExtensionSupported(extension string) bool {
	e := C.CString(extension)
	defer C.free(unsafe.Pointer(e))

	r := C.glfwExtensionSupported(e)

	if r == C.GL_FALSE {
		return false
	}
	return true
}
