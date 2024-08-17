#version 330 core

uniform mat4 projection;
uniform mat4 model;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main()
{
  fragTexCoord = vertTexCoord;
  gl_Position = projection * model * vec4(vert, 1.0);
}
