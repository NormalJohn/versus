import { _decorator, Component, Node } from 'cc';
const { ccclass, property } = _decorator;

@ccclass('matchSceneScript')
export class matchSceneScript extends Component {
    start() {

    }
    switchToScene1() {
     cc.director.loadScene("matchScene");
    }

    switchToScene2() {
        cc.director.loadScene("characterScene");
    }

    update(deltaTime: number) {
        
    }
}

