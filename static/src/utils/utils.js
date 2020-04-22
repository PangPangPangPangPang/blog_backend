/**
 * Created by Max on 2019-03-28.
 */
function isPC() {
  const userAgentInfo = navigator.userAgent
  const Agents = ['Android', 'iPhone', 'SymbianOS', 'Windows Phone', 'iPod']
  let flag = true
  for (let v = 0; v < Agents.length; v += 1) {
    if (userAgentInfo.indexOf(Agents[v]) > 0) {
      flag = false
      break
    }
  }
  return flag
}

function getLocalTime(ts, fmt) {
  const date = new Date(ts * 1000)
  const o = {
    'M+': date.getMonth() + 1,
    'd+': date.getDate(),
    'h+': date.getHours(),
    'm+': date.getMinutes(),
    's+': date.getSeconds(),
    'q+': Math.floor((date.getMonth() + 3) / 3),
    S: date.getMilliseconds(),
  }
  let ret = fmt
  if (/(y+)/.test(fmt)) {
    ret = ret.replace(
      RegExp.$1,
      `${date.getFullYear()}`.substr(4 - RegExp.$1.length),
    )
  }
  Object.keys(o).forEach((k) => {
    if (new RegExp(`(${k})`).test(fmt)) {
      ret = ret.replace(
        RegExp.$1,
        RegExp.$1.length === 1 ? o[k] : `00${o[k]}`.substr(`${o[k]}`.length),
      )
    }
  })
  return ret
}

export { isPC, getLocalTime }
