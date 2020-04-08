let ctx = null;
const tileW = 40, tileH = 40;
const mapW = 20, mapH = 20;
let gameMap;

window.onload = function () {
    ctx = document.getElementById('game').getContext("2d");
    initGameMap();
    ctx.font = "bold 10pt sans-serif";
};

window.onbeforeunload = function () {
    ctx = document.getElementById('game').getContext("2d");
    initGameMap();
    ctx.font = "bold 10pt sans-serif";
    return null;
};

function initGameMap() {
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", `http://localhost:8090/init?time=${new Date().getTime()}`, false ); // false for synchronous request
    xmlHttp.send( null );
    let responseText = xmlHttp.responseText;
    console.debug(JSON.parse(responseText));

    return JSON.parse(responseText);
}

function getGameMap() {
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", "http://localhost:8090/next?time="+new Date().getTime(), false ); // false for synchronous request
    xmlHttp.send( null );
    let responseText = xmlHttp.responseText;
    console.debug(JSON.parse(responseText));

    return JSON.parse(responseText);
}

function drawGame() {
    if (ctx == null) { return; }

    gameMap = getGameMap();
    for (let y = 0; y < mapH; ++y) {
        for (let x = 0; x < mapW; ++x) {
            if (gameMap[((y * mapW) + x)] === 0) {
                ctx.fillStyle = "#000000";
            } else {
                ctx.fillStyle = "#ccffcc";
            }

            ctx.fillRect(x * tileW, y * tileH, tileW, tileH);
        }
    }
}

function myLoop () {
    setTimeout(function () {
            drawGame();
            myLoop();
        },
        300);
}

myLoop();