#version 330 core

uniform sampler2D tex;
uniform vec3 spriteColor;

in vec2 fragTexCoord;

out vec4 outputColor;

void main()
{
  outputColor = vec4(spriteColor, 1.0) * texture(tex, fragTexCoord);
}
