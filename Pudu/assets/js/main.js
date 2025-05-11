// Navbar scroll effect and logo name reveal
window.addEventListener('scroll', function () {
    const header = document.querySelector('.header');
    const logoName = document.getElementById('logo-name');
    header.classList.toggle('scrolled', window.scrollY > 50);
    logoName.style.opacity = window.scrollY > 50 ? '1' : '0';
});

// WhatsApp button hover effect
const whatsappBtn = document.querySelector('.whatsapp-btn');
whatsappBtn.addEventListener('mouseenter', () => {
    whatsappBtn.style.transform = 'scale(1.1)';
});
whatsappBtn.addEventListener('mouseleave', () => {
    whatsappBtn.style.transform = 'scale(1)';
});

// Contact form submit logic
document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('contactForm');
    const emailInput = document.getElementById('email');
    const emailError = document.getElementById('email-error');

    form.addEventListener('submit', function (e) {
        e.preventDefault();

        const name = document.getElementById('name').value.trim();
        const email = emailInput.value.trim();
        const subject = document.getElementById('subject').value.trim();
        const message = document.getElementById('message').value.trim();

        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
            emailError.style.display = 'block';
            return;
        } else {
            emailError.style.display = 'none';
        }

        // Solicitar el token de reCAPTCHA
        grecaptcha.ready(async function () {
            try {
                const token = await grecaptcha.execute('6Lc_KSsrAAAAABcGf9xq86NdiiXODc3yVt3oBAx9', { action: 'submit' });

                const payload = {
                    Nombre: name,
                    Email: email,
                    Asunto: subject,
                    Mensaje: message,
                    recaptchaToken: token
                };

                const response = await fetch('https://iu89buz27k.execute-api.us-east-1.amazonaws.com/prod/sendmail', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(payload)
                });

                const data = await response.json();

                if (data.result === 'Success') {
                    alert('Correo enviado correctamente. ¡Gracias por contactarnos!');
                    form.reset();
                } else {
                    alert('Error al enviar el correo.');
                    console.error(data);
                }
            } catch (err) {
                console.error('Error al ejecutar reCAPTCHA o enviar datos:', err);
                alert('Ocurrió un error al validar el reCAPTCHA.');
            }
        });
    });
});
