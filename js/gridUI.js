let ctx = null;
const tileW = 40, tileH = 40;
const mapW = 20, mapH = 20;
let currentGameMap;
let settings = { };
let shouldDrawGame = false;
let shouldDisplayAlert = false;

window.onload = function () {
    ctx = document.getElementById('game').getContext("2d");
    // initGameMap();
    ctx.font = "bold 10pt sans-serif";
};

window.onbeforeunload = function () {
    ctx = document.getElementById('game').getContext("2d");
    // initGameMap();
    ctx.font = "bold 10pt sans-serif";
    return null;
};

function submitSettings() {
    shouldDrawGame = false;
    let threshold = Number(document.getElementById("threshold").value);
    if (threshold == null || threshold < 0 || threshold > 100){
        alert("Must be between 0-100");
        return;
    }
    settings['threshold'] = threshold;
    initGameMap();
}



function initGameMap() {
    shouldDisplayAlert = true;
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if (xmlHttp.status == 200){
            let responseText = xmlHttp.response;
            currentGameMap = JSON.parse(responseText);
            shouldDrawGame = true;
            myLoop();
        }
    };
    xmlHttp.open( "POST", `http://localhost:8090/init?time=${new Date().getTime()}`, false ); // false for synchronous request
    xmlHttp.send(JSON.stringify(settings));
    // todo: maybe raise exception if error
}

function getGameMapNextStep() {
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", `http://localhost:8090/next?time=${new Date().getTime()}`, false ); // false for synchronous request
    xmlHttp.send( null );
    let responseText = xmlHttp.response;
    return JSON.parse(responseText);
}

function drawGame() {
    if (ctx == null) { return false; }

    for (let y = 0; y < mapH; ++y) {
        for (let x = 0; x < mapW; ++x) {
            if (currentGameMap[((y * mapW) + x)] === 0) {
                ctx.fillStyle = "#000000";
            } else {
                ctx.fillStyle = "#ccffcc";
            }

            ctx.fillRect(x * tileW, y * tileH, tileW, tileH);
        }
    }

    let newGameMap = getGameMapNextStep();
    if (JSON.stringify(currentGameMap) === JSON.stringify(newGameMap)) {
        shouldDrawGame = false;
    } else {
        currentGameMap = newGameMap;
        shouldDrawGame = true;
    }
}

function myLoop () {
    setTimeout(function () {
            if (shouldDrawGame) {
                drawGame();
                myLoop();
            }
            else {
                if (shouldDisplayAlert){
                    alert("Game finished!")
                    shouldDisplayAlert = false;
                }
            }
        },
        100);
}

