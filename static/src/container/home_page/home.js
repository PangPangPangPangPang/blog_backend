/**
 * Created by Max on 02/03/2017.
 */
import React from 'react'
import './home.css'
import '../about/about.css'
import Footer from '../footer/footer'

export default class Home extends React.Component {
  imgStyle = {
    height: document.body.clientHeight - 48,
    width: (document.body.clientHeight / 1000) * 839,
  };

  constructor(props) {
    super(props)
    this.state = { show: true }
  }

  componentDidMount() {
    this.renderConvas()
  }

  click = () => {
    const { show } = this.state
    this.setState({
      show: !show,
    })
  };

  loop = (ctx) => {
    const { canvas } = this
    canvas.width = canvas.parentNode.offsetWidth
    canvas.height = canvas.parentNode.offsetHeight
    ctx.fillStyle = 'rgba(0,222,255, 0.2)'
    ctx.beginPath()
    ctx.moveTo(0, canvas.height / 2)
    ctx.lineTo(canvas.width, canvas.height / 2)
    ctx.lineTo(canvas.width, canvas.height)
    ctx.lineTo(0, canvas.height)
    ctx.lineTo(0, canvas.height / 2)
    ctx.closePath()
    ctx.fill()
  };

  renderConvas = () => {
    const { canvas } = this
    const ctx = canvas.getContext('2d')
    window.requestAnimFrame = (() => window.requestAnimationFrame
      || window.webkitRequestAnimationFrame
      || window.mozRequestAnimationFrame
      || function cb(callback) {
        window.setTimeout(callback, 1000 / 60)
      })()
    let step = 0
    const lines = [
      'rgba(0, 0, 0, 0.4)',
      'rgba(0, 0, 0, 0.4)',
      'rgba(0, 0, 0, 0.4)',
    ]
    const loop = () => {
      ctx.clearRect(0, 0, canvas.width, canvas.height)
      step += 1
      for (let j = lines.length - 1; j >= 0; j -= 1) {
        ctx.fillStyle = lines[j]
        const angle = ((step + j * 45) * Math.PI) / 180
        const deltaHeight = Math.sin(angle) * 50
        const deltaHeightRight = Math.cos(angle) * 50
        ctx.beginPath()
        ctx.moveTo(0, canvas.height / 4 + deltaHeight)
        ctx.bezierCurveTo(
          canvas.width / 2,
          canvas.height / 4 + deltaHeight - 50,
          canvas.width / 2,
          canvas.height / 4 + deltaHeightRight - 50,
          canvas.width,
          canvas.height / 4 + deltaHeightRight,
        )
        ctx.lineTo(canvas.width, canvas.height)
        ctx.lineTo(0, canvas.height)
        ctx.lineTo(0, canvas.height / 4 + deltaHeight)
        ctx.closePath()
        ctx.fill()
        // draw text
        ctx.font = 'bold 30px Arial'
        ctx.textAlign = 'center'
        ctx.textBaseLine = 'middle'
        ctx.fillStyle = 'white'
        ctx.fillText('Max Wang', canvas.width / 2, canvas.height / 4)

        ctx.font = 'bold 40px Arial'
        ctx.textAlign = 'center'
        ctx.textBaseLine = 'middle'
        ctx.fillStyle = 'white'
        ctx.fillText('iOS Coder', canvas.width / 2, canvas.height / 4 + 50)
      }
      window.requestAnimFrame(loop)
    }
    loop()
  };

  render() {
    return (
      <div
        className="home-template"
        style={{ height: document.body.clientHeight - 48 }}
      >
        <canvas
          ref={(c) => {
            this.canvas = c
          }}
          width={`${document.body.clientWidth}`}
          height="1500"
        />
        <Footer color="white" />
      </div>
    )
  }
}
