package main

import (
	"fmt"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"io/ioutil"
	"log"
)

import "github.com/avbaula/Go-SOIL/SOIL"

var elem_buffer_data []uint = []uint{0, 1, 2, 3}

var vert_buffer_data []float32 = []float32{
	-0.7, -0.7,
	0.7, -0.7,
	-0.7, 0.7,
	0.7, 0.7}
var tex_buffer_data []float32 = []float32{
	0.0, 0.0,
	1.0, 0.0,
	0.0, 1.0,
	1.0, 1.0}

var (
	TextCoordID, VertPosID             gl.AttribLocation
	TextureID                          gl.UniformLocation
	TextBuffer, VertBuffer, ElemBuffer gl.Buffer
	ProgID                             gl.Program
	VAO                                gl.VertexArray
	Texture                            gl.Texture
)

var (
	nFrames int32
	lTime   float64
	Run     bool
)

func DestroyWindow() {

	glfw.CloseWindow()
	glfw.Terminate()

}

func KeyEvent() {

	if glfw.Key(glfw.KeyEsc) == 1 && glfw.WindowParam(glfw.Opened) == 1 ||
		glfw.WindowParam(glfw.Opened) != 1 {
		Run = false
	}

}

func Reshape(w, h int) {
	gl.Viewport(0, 0, w, h)
}

func LoadTextFile(f string) string {

	file, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}
	return string(file)
}

func InitGL() {

	if gl.Init() != 0 {
		log.Panic("Error: OpenGL init\n")
	}

	gl.ClearColor(0.2, 0.2, 0.2, 1.0)

	vshader := gl.CreateShader(gl.VERTEX_SHADER)
	vshader.Source(LoadTextFile("vs.glsl"))
	vshader.Compile()
	if vshader.Get(gl.COMPILE_STATUS) != 1 {
		panic("vertex shader: " + vshader.GetInfoLog())
	}

	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fshader.Source(LoadTextFile("fs.glsl"))
	fshader.Compile()
	if fshader.Get(gl.COMPILE_STATUS) != 1 {
		panic("fragment shader: " + fshader.GetInfoLog())
	}

	ProgID = gl.CreateProgram()
	ProgID.AttachShader(vshader)
	ProgID.AttachShader(fshader)

	ProgID.Link()

	ProgID.DetachShader(vshader)
	ProgID.DetachShader(fshader)

	ProgID.Use()

	VertPosID = ProgID.GetAttribLocation("aVertPos")
	TextCoordID = ProgID.GetAttribLocation("aTexCoord")
	TextureID = ProgID.GetUniformLocation("uTexture")

	VAO = gl.GenVertexArray()
	VAO.Bind()

	VertBuffer = gl.GenBuffer()
	VertBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(vert_buffer_data)*4,
		vert_buffer_data, gl.STATIC_DRAW)
	VertPosID.AttribPointer(2, gl.FLOAT, false, 0, nil)
	VertPosID.EnableArray()

	TextBuffer = gl.GenBuffer()
	TextBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(tex_buffer_data)*4,
		tex_buffer_data, gl.STATIC_DRAW)
	TextCoordID.AttribPointer(2, gl.FLOAT, false, 0, nil)
	TextCoordID.EnableArray()

	ElemBuffer = gl.GenBuffer()
	ElemBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(elem_buffer_data)*4,
		elem_buffer_data, gl.STATIC_DRAW)

	gl.ProgramUnuse()
	VertBuffer.Bind(0)

	fmt.Println(SOIL.Last_result())

	Texture = gl.Texture(SOIL.Load_OGL_texture("image.png", SOIL.LOAD_AUTO,
		SOIL.CREATE_NEW_ID, SOIL.FLAG_INVERT_Y))

	Texture.Bind(gl.TEXTURE_2D)
	gl.ActiveTexture(gl.TEXTURE0)
	fmt.Println(SOIL.Last_result())
}

func Draw() {

	gl.Clear(gl.COLOR_BUFFER_BIT)

	ProgID.Use()

	VAO.Bind()

	TextureID.Uniform1i(0)

	gl.DrawElements(gl.TRIANGLE_STRIP, 4, gl.UNSIGNED_INT, nil)

	gl.ProgramUnuse()
	glfw.SwapBuffers()

}

func DestroyGL() {

	Texture.Delete()
	VertBuffer.Delete()
	TextBuffer.Delete()
	ElemBuffer.Delete()
	VertPosID.DisableArray()
	TextCoordID.DisableArray()
}

func main() {

	if err := glfw.Init(); err != nil {
		log.Panic("glfw Error:", err)
	}

	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 2)
	glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	glfw.OpenWindow(800, 600, 0, 0, 0, 0, 0, 0, glfw.Windowed)

	glfw.SetWindowTitle("Texture")
	glfw.SetWindowSizeCallback(Reshape)

	InitGL()

	Run = true
	for Run {
		Draw()
		KeyEvent()
	}

	DestroyGL()
	defer DestroyWindow()

}
