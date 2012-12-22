/*
 Go bindings for libSOIL
*/
package SOIL

//#cgo  LDFLAGS:  -lSOIL -lGL -lm
//#include <SOIL/SOIL.h>
import "C"
import "unsafe"

const (
	LOAD_AUTO              = C.SOIL_LOAD_AUTO
	LOAD_L                 = C.SOIL_LOAD_L
	LOAD_LA                = C.SOIL_LOAD_LA
	LOAD_RGB               = C.SOIL_LOAD_RGB
	LOAD_RGBA              = C.SOIL_LOAD_RGBA
	CREATE_NEW_ID          = C.SOIL_CREATE_NEW_ID
	FLAG_POWER_OF_TWO      = C.SOIL_FLAG_POWER_OF_TWO
	FLAG_MIPMAPS           = C.SOIL_FLAG_MIPMAPS
	FLAG_TEXTURE_REPEATS   = C.SOIL_FLAG_TEXTURE_REPEATS
	FLAG_MULTIPLY_ALPHA    = C.SOIL_FLAG_MULTIPLY_ALPHA
	FLAG_INVERT_Y          = C.SOIL_FLAG_INVERT_Y
	FLAG_COMPRESS_TO_DXT   = C.SOIL_FLAG_COMPRESS_TO_DXT
	FLAG_DDS_LOAD_DIRECT   = C.SOIL_FLAG_DDS_LOAD_DIRECT
	FLAG_NTSC_SAFE_RGB     = C.SOIL_FLAG_NTSC_SAFE_RGB
	FLAG_CoCg_Y            = C.SOIL_FLAG_CoCg_Y
	FLAG_TEXTURE_RECTANGLE = C.SOIL_FLAG_TEXTURE_RECTANGLE
	SAVE_TYPE_TGA          = C.SOIL_SAVE_TYPE_TGA
	SAVE_TYPE_BMP          = C.SOIL_SAVE_TYPE_BMP
	SAVE_TYPE_DDS          = C.SOIL_SAVE_TYPE_DDS
	HDR_RGBE               = C.SOIL_HDR_RGBE
	HDR_RGBdivA            = C.SOIL_HDR_RGBdivA
	HDR_RGBdivA2           = C.SOIL_HDR_RGBdivA2
)

func Load_OGL_texture(filename string,
	force_channels int32,
	reuse_texture_ID uint32,
	flags uint32) uint {
	return uint(C.SOIL_load_OGL_texture(C.CString(filename),
		C.int(force_channels),
		C.uint(reuse_texture_ID),
		C.uint(flags)))
}

func Load_OGL_cubmap(x_pos_file, x_nerg_file string,
	y_pos_file, y_nerg_file string,
	z_pos_file, z_neg_file string,
	force_channels int32,
	reuse_texture_ID uint,
	flags uint) uint {
	return uint(C.SOIL_load_OGL_cubemap(C.CString(x_pos_file), C.CString(x_nerg_file),
		C.CString(y_pos_file), C.CString(y_nerg_file),
		C.CString(z_pos_file), C.CString(z_neg_file),
		C.int(force_channels),
		C.uint(reuse_texture_ID),
		C.uint(flags)))
}

func Load_OGL_single_cubmap(filename string,
	face_order [6]string,
	force_channels int32,
	reuse_texture_ID uint,
	flags uint) uint {
	if len(face_order) == 0 {
		panic("Invalid param face_order:  the order of the faces in the file, any combination of NSWEUD, for North, South, Up, etc.")
	}
	return uint(C.SOIL_load_OGL_single_cubemap(C.CString(filename),
		C.CString(face_order[0]),
		C.int(force_channels),
		C.uint(reuse_texture_ID),
		C.uint(flags)))
}

func Load_OGL_HDR_texture(filename string,
	fake_HDR_format int32,
	rescale_to_max int32,
	reuse_texture_ID uint,
	flags uint) uint {
	return uint(C.SOIL_load_OGL_HDR_texture(C.CString(filename),
		C.int(fake_HDR_format),
		C.int(rescale_to_max),
		C.uint(reuse_texture_ID),
		C.uint(flags)))
}

func Load_OGL_texture_from_memory(buffer *byte,
	buffer_length int32,
	force_channels int32,
	reuse_texture_ID uint,
	flags uint) uint {
	return uint(C.SOIL_load_OGL_texture_from_memory((*C.uchar)(unsafe.Pointer(buffer)),
		C.int(buffer_length),
		C.int(force_channels),
		C.uint(reuse_texture_ID),
		C.uint(flags)))
}

func Load_OGL_cubmap_from_memory(x_pos_buffer *byte, x_pos_buffer_length int32,
	x_neg_buffer *byte, x_neg_buffer_length int32,
	y_pos_buffer *byte, y_pos_buffer_length int32,
	y_neg_buffer *byte, y_neg_buffer_length int32,
	z_pos_buffer *byte, z_pos_buffer_length int32,
	z_neg_buffer *byte, z_neg_buffer_length int32,
	force_channels int32,
	reuse_texture_ID uint,
	flags uint) uint {
	return uint(C.SOIL_load_OGL_cubemap_from_memory((*C.uchar)(unsafe.Pointer(x_pos_buffer)), C.int(x_pos_buffer_length),
		(*C.uchar)(unsafe.Pointer(x_neg_buffer)), C.int(x_neg_buffer_length),
		(*C.uchar)(unsafe.Pointer(y_pos_buffer)), C.int(y_pos_buffer_length),
		(*C.uchar)(unsafe.Pointer(y_neg_buffer)), C.int(y_neg_buffer_length),
		(*C.uchar)(unsafe.Pointer(z_pos_buffer)), C.int(z_pos_buffer_length),
		(*C.uchar)(unsafe.Pointer(z_neg_buffer)), C.int(z_neg_buffer_length),
		C.int(force_channels),
		C.uint(reuse_texture_ID),
		C.uint(flags)))
}

func Load_OGL_single_cubemap_from_memory(buffer *byte,
	buffer_length int32,
	face_order [6]string,
	force_channels int32,
	reuse_texture_ID uint,
	flags uint) uint {
	return uint(C.SOIL_load_OGL_single_cubemap_from_memory((*C.uchar)(unsafe.Pointer(buffer)),
		C.int(buffer_length),
		C.CString(face_order[0]),
		C.int(force_channels),
		C.uint(reuse_texture_ID),
		C.uint(flags)))
}

func Create_OGL_texture(data *byte, width int32, height int32, channels int32,
	reuse_texture_ID uint8,
	flags uint8) uint {
	return uint(C.SOIL_create_OGL_texture((*C.uchar)(unsafe.Pointer(data)),
		C.int(width),
		C.int(height),
		C.int(channels), C.uint(reuse_texture_ID), C.uint(flags)))
}

func Create_OGL_single_cubemap(data *byte,
	width int32, height int32,
	channels int32,
	face_order [6]string,
	reuse_texture_ID uint,
	flags uint) uint {
	return uint(C.SOIL_create_OGL_single_cubemap((*C.uchar)(unsafe.Pointer(data)),
		C.int(width),
		C.int(height),
		C.int(channels),
		(*C.char)(unsafe.Pointer(&face_order[0])),
		C.uint(reuse_texture_ID),
		C.uint(flags)))
}

func Save_screenshot(filename string,
	image_type int32,
	x int32, y int32,
	width int32, height int32) int {
	return int(C.SOIL_save_screenshot(C.CString(filename),
		C.int(image_type),
		C.int(x), C.int(y),
		C.int(width), C.int(height)))
}

func Load_image(filename string,
	width *int32, height *int32,
	channels *int32, force_channels int32) *byte {
	return (*byte)(C.SOIL_load_image(C.CString(filename),
		(*C.int)(unsafe.Pointer(width)),
		(*C.int)(unsafe.Pointer(height)),
		(*C.int)(unsafe.Pointer(channels)),
		C.int(force_channels)))
}

func Load_image_from_memory(buffer *byte, buffer_length int32,
	width *int32, height *int32,
	channels *int32, force_channels int32) *byte {
	return (*byte)(C.SOIL_load_image_from_memory((*C.uchar)(unsafe.Pointer(buffer)),
		C.int(buffer_length),
		(*C.int)(unsafe.Pointer(width)),
		(*C.int)(unsafe.Pointer(height)),
		(*C.int)(unsafe.Pointer(channels)), C.int(force_channels)))
}

func Save_image(filename string,
	image_type int32,
	width int32, height int32,
	channels int32, data *byte) int {
	return int(C.SOIL_save_image(C.CString(filename),
		C.int(image_type),
		C.int(width),
		C.int(height),
		C.int(channels),
		(*C.uchar)(unsafe.Pointer(data))))
}

func Free_image_data(t *byte) {
	C.SOIL_free_image_data((*C.uchar)(unsafe.Pointer(t)))
}

func Last_result() string {
	return C.GoString(C.SOIL_last_result())
}
