#version 400 core

in vec2 pass_textureCoords;
in vec3 surfaceNormal;
in vec3 toLightVector;

out vec4 outColour;

uniform sampler2D textureSampler;
uniform vec3 lightColour;

void main(void){
    // we need to normalize for dot product
    vec3 unitNormal = normalize(surfaceNormal);
    vec3 unitLight = normalize(toLightVector);

    float lightIntensity = dot(unitNormal, unitLight);
    float brightness = max(lightIntensity, 0.0); // don't want negative intensity
    vec3 diffuseLight = brightness * lightColour; // our light should be coloured

    outColour = vec4(diffuseLight, 1.0) * texture(textureSampler, pass_textureCoords);
}
