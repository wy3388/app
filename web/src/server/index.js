import {spawn} from 'child_process'
import iconv from 'iconv-lite'

const startServer = () => {
    const isDevelopment = process.env.NODE_ENV !== 'production'
    const path = isDevelopment ? './' : './resources/'
    let cmd = ''
    if (process.platform === 'win32') {
        cmd = path + 'server.exe'
    }
    if (process.platform === 'linux' || process.platform === 'darwin') {
        cmd = path + 'server'
    }
    if (cmd === '') {
        console.log('后台服务启动失败')
        return
    }
    const s = spawn(cmd)
    s.stdout.on('data', data => {
        if (process.platform === 'win32') {
            console.log(iconv.decode(Buffer.from(data), 'gbk'));
            return
        }
        console.log(data)
    })
    s.on('close', () => {
        console.log('后台服务退出成功')
    })
}

const stopServer = () => {
    let cmd = ''
    let args = []
    if (process.platform === 'win32') {
        cmd = 'taskkill'
        args = ['/f', '/t', '/im', 'server.exe']
    }
    if (cmd === '') {
        console.log('停止服务参数错误')
        return
    }
    const s = spawn(cmd, args)
    s.stdout.on('data', data => {
        if (process.platform === 'win32') {
            console.log(iconv.decode(Buffer.from(data), 'gbk'));
            return
        }
        console.log(data)
    })
}

export {
    startServer,
    stopServer
}
