const signUpButton = document.getElementById('signUp');
const signInButton = document.getElementById('signIn');
const container = document.getElementById('container');

signUpButton.addEventListener('click', () => {
	container.classList.add("right-panel-active");
});

signInButton.addEventListener('click', () => {
	container.classList.remove("right-panel-active");
});

function getQueryParam(param) {
  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);
  return urlParams.get(param);
}

const errorValue = getQueryParam('error');
const usernameErrorDiv = document.getElementById('usernameError');
const passwordErrorDiv = document.getElementById('passwordError'); 
const usernameErrorSignDiv = document.getElementById('usernameErrorSignUp');
const mailErrorDiv = document.getElementById('mailError');

if (errorValue === 'badname') {
  console.log('Nom d\'utilisateur incorrect');
  usernameErrorDiv.textContent = 'Nom d\'utilisateur incorrect'; 
  usernameErrorDiv.style.color = 'red';
  usernameErrorDiv.style.fontSize = '11px'; 
} else if (errorValue === 'badpassword') {
  console.log('Mot de passe incorrect');
  passwordErrorDiv.textContent = 'Mot de passe incorrect'; 
  passwordErrorDiv.style.color = 'red';
  passwordErrorDiv.style.fontSize = '11px'; 
} else if (errorValue === 'badnameSignUp') {
  console.log('Nom d\'utilisateur déjà utilisé');
  usernameErrorSignDiv.textContent = 'Nom d\'utilisateur déjà utilisé';
  usernameErrorSignDiv.style.color = 'red';
  usernameErrorSignDiv.style.fontSize = '11px';
  container.classList.add("right-panel-active");
} else if (errorValue === 'badmailSignUp') {
  console.log('Email déjà utilisé');
  mailErrorDiv.textContent = 'Email déjà utilisé';
  mailErrorDiv.style.color = 'red';
  mailErrorDiv.style.fontSize = '11px';
  container.classList.add("right-panel-active");
}

  
