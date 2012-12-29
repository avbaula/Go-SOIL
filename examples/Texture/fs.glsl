#version 150 core
in vec2 vTexCoord;
uniform sampler2D uTexture;
void main()
{
	gl_FragColor = texture2D(uTexture, vTexCoord);
} 
