const fetchPODInfo = async () => {
    try {
        let res = await fetch('/pod');
        return await res.json();
    } catch (err) {
        console.error(err);
    }
};

const updatePODInfo = async () => {
    let pod = await fetchPODInfo();
    if (!pod) {
        pod = {hostname: "---"}
    }
    let elem = document.querySelector('#pod-hostname');
    if (elem.textContent !== pod.hostname) {
        elem.parentElement.classList.remove('bounceIn');
        elem.parentElement.classList.remove('animated');
        setTimeout(() => {
            elem.parentElement.classList.add('animated');
            elem.parentElement.classList.add('bounceIn');
            elem.textContent = pod.hostname;
        }, 20);
    }
};

const runProgress = (totalTime) => {
    let target = totalTime; // assume ms
    let incr = 15; //ms
    let current = 0;
    return setInterval(() => {
        current += incr;
        if (current > target) {
            current = 0;
        }
        let percentage = (current * 100) / target;
        document.querySelector('#bottom-bar').style.width = `${percentage}%`
    }, incr);
};


const loopFetch = () => {
    let interval = 10 * 1000;
    setInterval(() => {
        updatePODInfo();
    }, interval);
    runProgress(interval);
};
updatePODInfo();
loopFetch();