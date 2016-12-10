package render

import (
	"github.com/DrTeePot/game/model"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func Prepare() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// TODO RawModel as an interface?
func Render(model model.RawModel) {
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, model.BufferID())

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	// gl.DrawElements(gl.TRIANGLES, model.VertexCount(), gl.UNSIGNED_INT, nil)
	gl.DrawArrays(gl.TRIANGLES, 0, model.VertexCount())

	gl.DisableVertexAttribArray(0)
}
