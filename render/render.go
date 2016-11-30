package render

import (
	"github.com/DrTeePot/game/model"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func Prepare() {
	gl.ClearColor(1, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// TODO RawModel as an interface?
func Render(model model.RawModel) {
	gl.BindVertexArray(model.VAOID())
	gl.EnableVertexAttribArray(0)
	gl.DrawArrays(gl.TRIANGLES, 0, model.VertexCount())
	gl.DisableVertexAttribArray(0)
	gl.BindVertexArray(0)
}
