
const zid =  document.getElementById('zidd');
zid.addEventListener('click', () => {
    zid.style.display = 'flex'; 
    document.getElementById('zid').classList.add('post')
    document.getElementById('zid').classList.remove('post')
})

const auth = document.getElementById('auth');
const sowboton = document.getElementById('show')


sowboton.addEventListener('click', () => {
    auth.style.display = 'flex';
    document.getElementById('page-logen').classList.add('active');
    document.getElementById('cection-page').classList.remove('active');
});

function togglePage() {
    const loginPage = document.getElementById('page-logen');
    const seconnecterPage = document.getElementById('cection-page');
  
    loginPage.classList.toggle('active');
    seconnecterPage.classList.toggle('active');
  }

function updateClock() {
    const hours = document.getElementById('hrs');
    const minutes = document.getElementById('mins');
    const seconds = document.getElementById('secs');

    const hhCircle = document.getElementById('hh');
    const mmCircle = document.getElementById('mm');
    const ssCircle = document.getElementById('ss');

    const now = new Date();
    let hrs = now.getHours();
    let mins = now.getMinutes();
    let secs = now.getSeconds();

    hrs = hrs % 12 || 12;

    hours.innerText = hrs.toString().padStart(2, '0');
    minutes.innerText = mins.toString().padStart(2, '0');
    seconds.innerText = secs.toString().padStart(2, '0');

    const circumference = 251.2;

    hhCircle.style.strokeDashoffset = circumference - (circumference * hrs) / 12;
    mmCircle.style.strokeDashoffset = circumference - (circumference * mins) / 60;
    ssCircle.style.strokeDashoffset = circumference - (circumference * secs) / 60;
}

setInterval(updateClock, 1000);
updateClock();

// r::lengthj