#version 400 core

in vec2 pass_textureCoords;
in vec3 surfaceNormal;
in vec3 toLightVector;
in vec3 toCameraVector;

out vec4 outColour;

uniform sampler2D textureSampler;
uniform vec3 lightColour;
uniform float shineDamper;
uniform float reflectivity;

void main(void){
    // we need to normalize for dot product
    vec3 unitNormal = normalize(surfaceNormal);
    vec3 unitLight = normalize(toLightVector);

    float lightIntensity = dot(unitNormal, unitLight);
    // 0.1 is ambient lighting
    float brightness = max(lightIntensity, 0.1); // don't want negative intensity
    vec3 diffuseLight = brightness * lightColour; // our light should be coloured

    // specular
    vec3 unitCamera = normalize(toCameraVector);
    vec3 lightDirection = -unitLight;
    vec3 reflectedLight = reflect(lightDirection, unitNormal);

    float specularFactor = dot(reflectedLight, unitCamera);
    specularFactor = max(specularFactor, 0.0);
    float dampedFactor = pow(specularFactor, shineDamper);
    vec3 finalSpecular = dampedFactor * reflectivity * lightColour;

    outColour = vec4(diffuseLight, 1.0) * texture(textureSampler, pass_textureCoords) + vec4(finalSpecular, 1.0);
}
