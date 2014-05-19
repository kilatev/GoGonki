renderer = new PIXI.CanvasRenderer(350,250)
dX=4
dY=4
document.body.appendChild(renderer.view)
stage = new PIXI.Stage
logoTexture = PIXI.Texture.fromImage("http://www.javascriptoo.com/application/assets/images/favicon.png")
logo = new PIXI.Sprite(logoTexture)
logo.position.x = 88
logo.position.y = 10
stage.addChild(logo)
requestAnimationFrame(animate)

animate = ->
    //logo.rotation += 0.05
    if (logo.position.y <= 0 || logo.position.y >= 202)
      dY *= -1
    if (logo.position.x <= 0 || logo.position.x >= 302)
      dX *= -1
    logo.position.x += dX
    logo.position.y += dY
    renderer.render(stage)
    requestAnimationFrame(animate)
