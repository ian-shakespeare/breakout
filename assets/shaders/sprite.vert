#version 330 core

uniform mat4 projection;
uniform mat4 model;

in vec2 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main()
{
  fragTexCoord = vertTexCoord;
  gl_Position = projection * model * vec4(vert.xy, 0.0, 1.0);
}
