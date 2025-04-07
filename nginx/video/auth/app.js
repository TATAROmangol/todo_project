document.addEventListener('DOMContentLoaded', function() {
    const API_BASE = '/api/auth';
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');
    const messageDiv = document.getElementById('message');
    const tabBtns = document.querySelectorAll('.tab-btn');
    const tabContents = document.querySelectorAll('.tab-content');

    // Tab switching
    tabBtns.forEach(btn => {
        btn.addEventListener('click', () => {
            const tabId = btn.getAttribute('data-tab');
            
            tabBtns.forEach(b => b.classList.remove('active'));
            tabContents.forEach(c => c.classList.remove('active'));
            
            btn.classList.add('active');
            document.getElementById(tabId).classList.add('active');
        });
    });

    // Login handler
    loginForm.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const username = document.getElementById('loginName').value;
        const password = document.getElementById('loginPassword').value;
        
        try {
            const response = await fetch(`${API_BASE}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name: username,
                    password: password
                }),
                credentials: 'include' // Для работы с cookies
            });

            if (!response.ok) {
                throw new Error('Login failed');
            }

            showMessage('Login successful! Redirecting...', 'success');
            
            // Переход на сервис задач через 1 секунду
            setTimeout(() => {
                window.location.href = '/todo/';
            }, 1000);
            
        } catch (error) {
            showMessage(error.message, 'error');
        }
    });

    // Register handler
    registerForm.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const username = document.getElementById('registerName').value;
        const password = document.getElementById('registerPassword').value;
        
        try {
            const response = await fetch(`${API_BASE}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name: username,
                    password: password
                }),
                credentials: 'include' // Для работы с cookies
            });

            if (!response.ok) {
                throw new Error('Registration failed');
            }

            showMessage('Registration successful! Redirecting...', 'success');
            
            // Переход на сервис задач через 1 секунду
            setTimeout(() => {
                window.location.href = '/todo/';
            }, 1000);
            
        } catch (error) {
            showMessage(error.message, 'error');
        }
    });

    function showMessage(message, type) {
        messageDiv.textContent = message;
        messageDiv.className = `message ${type}`;
        
        setTimeout(() => {
            messageDiv.textContent = '';
            messageDiv.className = 'message';
        }, 3000);
    }
});