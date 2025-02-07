const zid = document.getElementById('zidpost'); // page zid post
const add = document.getElementById('add'); // boton +
const auth = document.getElementById('auth'); // page kamla li fiha log in w se ..
const sowboton = document.getElementById('show') // butun se conecter
const loginPage = document.getElementById('page-logen'); // page login lwla
const seconnecterPage = document.getElementById('cection-page'); // page login tanya
const registerForm = document.getElementById("registerForm");

registerForm.addEventListener("submit", async (event) => {
    event.preventDefault();

    const formData = new FormData(registerForm);
    const data = {
        username: formData.get("username"),
        email: formData.get("email"),
        password: formData.get("password"),
    };

    try {
        const response = await fetch("/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });

        if (response.ok) {
            alert("User registered successfully!");
        } else {
            const error = await response.text();
            alert(`Error: ${error}`);
        }
    } catch (error) {
        console.error("Error:", error);
        alert("An error occurred during registration.");
    }
    console.log(formData.username)
    console.log(formData.email)
    console.log(formData.password)
});


add.addEventListener('click', () => { // byan add le post
    zid.style.display = 'flex';
    auth.style.display = 'none'
    document.getElementById('zidpost').classList.add('active')
})

sowboton.addEventListener('click', () => { 
    auth.style.display = 'flex';
    zid.style.display = 'none';
    document.getElementById('page-logen').classList.add('active');
    document.getElementById('cection-page').classList.remove('active');
});

function togglePage() {
  
    loginPage.classList.toggle('active');
    seconnecterPage.classList.toggle('active');
  }


// page heders

// function updateClock() {
//     const hours = document.getElementById('hrs');
//     const minutes = document.getElementById('mins');
//     const seconds = document.getElementById('secs');

//     const hhCircle = document.getElementById('hh');
//     const mmCircle = document.getElementById('mm');
//     const ssCircle = document.getElementById('ss');

//     const now = new Date();
//     let hrs = now.getHours();
//     let mins = now.getMinutes();
//     let secs = now.getSeconds();

//     hrs = hrs % 12 || 12;

//     hours.innerText = hrs.toString().padStart(2, '0');
//     minutes.innerText = mins.toString().padStart(2, '0');
//     seconds.innerText = secs.toString().padStart(2, '0');

//     const circumference = 251.2;

//     hhCircle.style.strokeDashoffset = circumference - (circumference * hrs) / 12;
//     mmCircle.style.strokeDashoffset = circumference - (circumference * mins) / 60;
//     ssCircle.style.strokeDashoffset = circumference - (circumference * secs) / 60;
// }

// setInterval(updateClock, 1000);
// updateClock();
// // mergen
// //dd