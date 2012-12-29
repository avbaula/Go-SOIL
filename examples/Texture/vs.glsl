#version 150 core
attribute vec2 aVertPos;
attribute vec2 aTexCoord;
out vec2 vTexCoord;
 void main()
{
	vTexCoord = aTexCoord;
	gl_Position = vec4(aVertPos, 1.0f, 1.0f);
} 
