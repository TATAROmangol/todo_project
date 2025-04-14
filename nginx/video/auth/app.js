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
            clearMessage();
        });
    });

    // Login handler
    loginForm.addEventListener('submit', async function(e) {
        e.preventDefault();
        clearMessage();
        
        const username = document.getElementById('loginName').value.trim();
        const password = document.getElementById('loginPassword').value;
        
        if (!username || !password) {
            showMessage('Please fill in all fields', 'error');
            return;
        }
        
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
                credentials: 'include'
            });
            
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            showMessage('Login successful! Redirecting...', 'success');
            
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
        clearMessage();
        
        const username = document.getElementById('registerName').value.trim();
        const password = document.getElementById('registerPassword').value;
        const confirmPassword = document.getElementById('registerConfirmPassword').value;
        
        // Validation
        if (!username || !password || !confirmPassword) {
            showMessage('Please fill in all fields', 'error');
            return;
        }
        
        if (password !== confirmPassword) {
            showMessage('Passwords do not match', 'error');
            return;
        }
        
        if (password.length < 6) {
            showMessage('Password must be at least 6 characters', 'error');
            return;
        }
        
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
                credentials: 'include'
            });
            
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            showMessage('Registration successful! Redirecting...', 'success');
            
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
    }
    
    function clearMessage() {
        messageDiv.textContent = '';
        messageDiv.className = 'message';
    }
});