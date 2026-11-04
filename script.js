const steps = [
  {
    number: '01',
    prompt: '“Remind me to check the garden lights<br />when the sun goes down.”',
    response: 'I’ll remind you at sunset today. No extra setup needed.'
  },
  {
    number: '02',
    prompt: '“What do I need to finish before<br />tomorrow morning?”',
    response: 'I found three open tasks and grouped them by priority.'
  },
  {
    number: '03',
    prompt: '“Save this article and give me<br />the short version.”',
    response: 'Saved. Here is the useful part in three clear points.'
  },
  {
    number: '04',
    prompt: '“Done. What did you change<br />and what should I know?”',
    response: 'Everything is ready. Here is a simple record of what happened.'
  }
];

const buttons = document.querySelectorAll('.step');
const index = document.querySelector('.demo-index strong');
const prompt = document.querySelector('.demo-prompt');
const response = document.querySelector('.demo-response p');

buttons.forEach((button) => {
  button.addEventListener('click', () => {
    const step = steps[Number(button.dataset.step)];
    buttons.forEach((item) => {
      item.classList.remove('active');
      item.setAttribute('aria-selected', 'false');
    });
    button.classList.add('active');
    button.setAttribute('aria-selected', 'true');
    index.textContent = step.number;
    prompt.innerHTML = step.prompt;
    response.textContent = step.response;
  });
});

const waitlistForm = document.querySelector('#waitlist-form');
const waitlistEmail = document.querySelector('#waitlist-email');
const waitlistCount = document.querySelector('#waitlist-count');
const waitlistStatus = document.querySelector('#waitlist-status');
let waitlistTotal = 1097;

const renderWaitlistCount = () => {
  waitlistCount.textContent = waitlistTotal.toLocaleString();
};

setInterval(() => {
  waitlistTotal = Math.max(1097, waitlistTotal + (Math.random() > 0.5 ? 1 : -1));
  renderWaitlistCount();
}, 15000);

waitlistForm.addEventListener('submit', (event) => {
  event.preventDefault();
  const email = waitlistEmail.value.trim().toLowerCase();

  if (!email || !waitlistEmail.checkValidity()) {
    waitlistStatus.textContent = 'Enter a valid email address.';
    return;
  }

  const joinedEmails = JSON.parse(localStorage.getItem('orion-waitlist-emails') || '[]');
  if (joinedEmails.includes(email)) {
    waitlistStatus.textContent = 'You are already on the list.';
    return;
  }

  joinedEmails.push(email);
  localStorage.setItem('orion-waitlist-emails', JSON.stringify(joinedEmails));
  waitlistTotal += 1;
  renderWaitlistCount();
  waitlistStatus.textContent = 'You are on the list. Welcome to Orion.';
  waitlistForm.reset();
});
