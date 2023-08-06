import { _decorator, Component, Node,EditBox,sys,director } from 'cc';
const { ccclass, property } = _decorator;
import { SERVER_URL } from './GlobalConfig';
@ccclass('loginPageScript')
export class loginPageScript extends Component {
    start() {

    }
    login(): void {
        console.log("clicked")
        const node  = this.node;
        const username = this.node.getChildByName('usernameInput').getComponent(EditBox).string;
        const password = this.node.getChildByName('passwordInput').getComponent(EditBox).string;
        if (username === "" || password === "") {
            alert("用户名或密码不能为空");
            return;
        }
        this.sendLoginRequest(username, password);
    }

    sendLoginRequest(username: string, password: string): void {
        const xhr = new XMLHttpRequest();
        xhr.open("POST", SERVER_URL + "/login");
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.onreadystatechange = () => {
            if (xhr.readyState === 4 && xhr.status === 200) {
                const response = JSON.parse(xhr.responseText);
                console.log(response);
                if (response.code === 200) {
                    sys.localStorage.setItem("token", response.token);
                    sys.localStorage.setItem("expire", response.expire);
                    sys.localStorage.setItem("username", username);
                    alert("登录成功");
                    director.loadScene("matchScene");
                } else {
                    alert("登录失败");
                }
            }
        };
        xhr.onerror = () => {
            alert("网络错误");
        };
        xhr.send(JSON.stringify({
            username,
            password,
        }));
    }


    update(deltaTime: number) {

    }
}

