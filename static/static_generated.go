// Code generated by generate_static.go; DO NOT EDIT.

package static

import "time"

var ModTime = time.Unix(0, 1602769004026352000)

const IndexHtml = "<!DOCTYPE html>\n<meta name=viewport content=width=device-width>\n<style>\n\nhtml {\n  font: 48px system-ui;\n  font-variant-numeric: tabular-nums;\n  background: #222;\n  color: white;\n  height: 100%;\n  display: flex;\n  align-items: center;\n  justify-content: center;\n}\n\nhtml.tick {\n  background: white;\n  color: black;\n}\n\n</style>\n<h1 id=clockTime></h1>\n<script type=module>\n\"use strict\";\n\nlet ticking = false;\nlet clockOffset;\nlet bestRtt;\n\nconst tick = now => {\n  if (clockOffset)\n    now -= clockOffset;\n  if (bestRtt)\n    now += bestRtt / 2;\n  const frame = Math.floor(now / (1000/60)) % 10000;\n  clockTime.textContent = frame;\n  document.documentElement.classList[Math.floor(frame / 100) % 2 ? 'add' : 'remove']('tick');\n  requestAnimationFrame(tick);\n};\n\nconst connectWs = reloadOnOpen => {\n  let pingInterval;\n  const ping = () => {\n    ws.send(JSON.stringify({\n      type: \"ping\",\n      startTime: Math.round(performance.now()),\n    }));\n  };\n  const ws = new WebSocket(`${location.protocol == 'https:' ? 'wss' : 'ws'}://${location.host}/ws`);\n  ws.onopen = e => {\n    if (reloadOnOpen)\n      location.reload(true);\n    pingInterval = setInterval(ping, 1000);\n    ping();\n  };\n  ws.onclose = e => {\n    clearInterval(pingInterval);\n    setTimeout(() => {\n      connectWs(true);\n    }, 1000);\n  };\n  ws.onmessage = e => {\n    const message = JSON.parse(e.data);\n    const { type } = message;\n    if (type == 'pong') {\n      const { startTime, serverTime } = message;\n      const now = performance.now();\n      const rtt = now - startTime;\n      bestRtt = bestRtt ? Math.min(bestRtt, rtt) : rtt;\n      clockOffset = clockOffset ? Math.min(clockOffset, now - serverTime) : now - serverTime;\n      if (!ticking) {\n        tick(performance.now());\n        ticking = true;\n      }\n    }\n  };\n};\nconnectWs();\n\n</script>\n"