let ctx = null;
const tileW = 35, tileH = 35;
const mapW = 20, mapH = 20;
let currentGameMap;
let settings = { };
let isGameFinished = false;
let stopLoop = false;

window.onload = function () {
    ctx = document.getElementById('game').getContext("2d");
};

window.onbeforeunload = function () {
    ctx = document.getElementById('game').getContext("2d");
};

function nextStep() {
    if (isGameFinished) {
        setInfoWindowTextValue("Game Finished");
        setInfoWindowTextOpacity(1);
        return;
    }
    if (!stopLoop){
        pause();
    }
    drawGame();
}

function pause(){
    stopLoop = !stopLoop;
    let pauseButton = document.getElementById("pauseButton");
    pauseButton.value = stopLoop ? "Continue" : "Pause";
    if (!stopLoop){
        myLoop();
    }
}

function setInfoWindowTextOpacity(opacity) {
    document.getElementById('infoWindowText').style.opacity = opacity;
}

function setInfoWindowTextValue(text) {
    document.getElementById('infoWindowText').innerText = text;
}

function start() {
    initializeNewGame();
    getSettingsFromForm();
    sendInitRequest();

    function initializeNewGame() {
        setInfoWindowTextOpacity(0);
        setInfoWindowTextValue("");
        isGameFinished = false;
        stopLoop = false;
        let pauseButton = document.getElementById("pauseButton");
        pauseButton.value = "Pause";
    }

    function sendInitRequest() {
        try{
            let xmlHttp = new XMLHttpRequest();
            xmlHttp.onreadystatechange = function() {
                if (xmlHttp.status == 200){
                    setInfoWindowTextOpacity(0);
                    setInfoWindowTextValue("");
                    let responseText = xmlHttp.response;
                    currentGameMap = JSON.parse(responseText);
                    myLoop();
                }
                else {
                    setInfoWindowTextValue("Network error!");
                    setInfoWindowTextOpacity(1);
                }
            };
            xmlHttp.open( "POST", `http://localhost:8090/init?time=${new Date().getTime()}`, false ); // false for synchronous request
            xmlHttp.send(JSON.stringify(settings));
        }
        catch (e) {
            setInfoWindowTextValue("Network error!");
            setInfoWindowTextOpacity(1);
        }
        // todo: maybe raise exception if error show alert

    }

    function getSettingsFromForm() {
        let threshold = Number(document.getElementById("threshold").value);
        if (threshold == null || threshold < 0 || threshold > 100){
            alert("Must be between 0-100");
            return;
        }
        settings['threshold'] = threshold;

        let liveSlotMinLiveNeighboursToKeepAlive = Number(document.getElementById("liveSlotMinLiveNeighboursToKeepAlive").value);
        if (liveSlotMinLiveNeighboursToKeepAlive == null || liveSlotMinLiveNeighboursToKeepAlive < 0 || liveSlotMinLiveNeighboursToKeepAlive > 8){
            alert("Must be between 0-8");
            return;
        }
        settings['liveSlotMinLiveNeighboursToKeepAlive'] = liveSlotMinLiveNeighboursToKeepAlive;

        let liveSlotMaxLiveNeighboursToKeepAlive = Number(document.getElementById("liveSlotMaxLiveNeighboursToKeepAlive").value);
        if (liveSlotMaxLiveNeighboursToKeepAlive == null || liveSlotMaxLiveNeighboursToKeepAlive < 0 || liveSlotMinLiveNeighboursToKeepAlive > 8){
            alert("Must be between 0-8");
            return;
        }
        settings['liveSlotMaxLiveNeighboursToKeepAlive'] = liveSlotMaxLiveNeighboursToKeepAlive;

        let deadSlotMinLiveNeighboursToBringToLife = Number(document.getElementById("deadSlotMinLiveNeighboursToBringToLife").value);
        if (deadSlotMinLiveNeighboursToBringToLife == null || deadSlotMinLiveNeighboursToBringToLife < 0 || deadSlotMinLiveNeighboursToBringToLife > 8){
            alert("Must be between 0-8");
            return;
        }
        settings['deadSlotMinLiveNeighboursToBringToLife'] = deadSlotMinLiveNeighboursToBringToLife;

        let deadSlotMaxLiveNeighboursToBringToLife = Number(document.getElementById("deadSlotMaxLiveNeighboursToBringToLife").value);
        if (deadSlotMaxLiveNeighboursToBringToLife == null || deadSlotMaxLiveNeighboursToBringToLife < 0 || deadSlotMaxLiveNeighboursToBringToLife > 8){
            alert("Must be between 0-8");
            return;
        }
        settings['deadSlotMaxLiveNeighboursToBringToLife'] = deadSlotMaxLiveNeighboursToBringToLife;
    }
}


function getGameMapNextStep() {
    try{
        let xmlHttp = new XMLHttpRequest();
        xmlHttp.open( "GET", `http://localhost:8090/next?time=${new Date().getTime()}`, false ); // false for synchronous request
        xmlHttp.send( null );
        let responseText = xmlHttp.response;
        return JSON.parse(responseText);
    }
    catch (e) {
        setInfoWindowTextValue("Network error!");
        setInfoWindowTextOpacity(1);
    }
}

function drawGame() {
    if (ctx == null)
    {
        return false;
    }

    fillCanvasWithGameMap(currentGameMap);

    let newGameMap = getGameMapNextStep();
    if (JSON.stringify(currentGameMap) === JSON.stringify(newGameMap)) {
        isGameFinished = true;
    } else {
        currentGameMap = newGameMap;
    }

    function fillCanvasWithGameMap(gameMap) {
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
}

function myLoop () {
    setTimeout(function () {
            if (stopLoop){
                return;
            }
            if (isGameFinished) {
                stopLoop = true;
                setInfoWindowTextOpacity(1);
                setInfoWindowTextValue("Game Finished");
                return;
            }
            drawGame();
            myLoop();
        },
        100);
}

