document.addEventListener('DOMContentLoaded', function() {
    const API_BASE = '/api/todo';
    const taskInput = document.getElementById('taskInput');
    const addTaskBtn = document.getElementById('addTaskBtn');
    const taskList = document.getElementById('taskList');
    const logoutBtn = document.getElementById('logoutBtn');

    // Первым делом проверяем авторизацию через GET запрос
    checkAuthAndLoadTasks();

    // Обработчики событий
    addTaskBtn.addEventListener('click', addTask);
    taskInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') addTask();
    });
    logoutBtn.addEventListener('click', logout);

    // Основные функции
    async function checkAuthAndLoadTasks() {
        try {
            showMessage('Checking authorization...');
            
            // Первый запрос - проверка авторизации через GET
            const response = await fetch(`${API_BASE}/get`, {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                if (response.status === 401) {
                    showMessage('Please login first', 'error');
                    setTimeout(() => window.location.href = '/auth/', 2000);
                    return;
                }
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            // Если авторизация успешна, активируем интерфейс
            taskInput.disabled = false;
            addTaskBtn.disabled = false;
            
            const tasks = await response.json();
            renderTasks(tasks);
        } catch (error) {
            console.error('Error:', error);
            showMessage(`Error: ${error.message}`, 'error');
            setTimeout(() => window.location.href = '/auth/', 2000);
        }
    }

    async function addTask() {
        const taskName = taskInput.value.trim();
        if (!taskName) {
            showMessage('Task name cannot be empty', 'error');
            return;
        }

        try {
            const response = await fetch(`${API_BASE}/post`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ name: taskName })
            });

            if (!response.ok) {
                if (response.status === 401) {
                    showMessage('Session expired, please login again', 'error');
                    setTimeout(() => window.location.href = '/auth/', 2000);
                    return;
                }
                throw new Error('Failed to add task');
            }

            taskInput.value = '';
            checkAuthAndLoadTasks(); // Перезагружаем список задач
            showMessage('Task added successfully!', 'success');
        } catch (error) {
            console.error('Error:', error);
            showMessage(`Error: ${error.message}`, 'error');
        }
    }

    async function deleteTask(taskId) {
        try {
            const response = await fetch(`${API_BASE}/delete`, {
                method: 'DELETE',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ id: taskId })
            });

            if (!response.ok) {
                if (response.status === 401) {
                    showMessage('Session expired, please login again', 'error');
                    setTimeout(() => window.location.href = '/auth/', 2000);
                    return;
                }
                throw new Error('Failed to delete task');
            }

            checkAuthAndLoadTasks(); // Перезагружаем список задач
            showMessage('Task deleted successfully!', 'success');
        } catch (error) {
            console.error('Error:', error);
            showMessage(`Error: ${error.message}`, 'error');
        }
    }

    async function logout() {
        try {
            const response = await fetch('/api/auth/logout', {
                method: 'POST',
                credentials: 'include'
            });

            if (!response.ok) {
                throw new Error('Logout failed');
            }

            window.location.href = '/auth/';
        } catch (error) {
            console.error('Error:', error);
            showMessage(`Error: ${error.message}`, 'error');
        }
    }

    function renderTasks(tasks) {
        if (!tasks || tasks.length === 0) {
            taskList.innerHTML = '<div class="loading">No tasks found. Add a new task!</div>';
            return;
        }

        taskList.innerHTML = '';
        tasks.forEach(task => {
            const taskElement = document.createElement('div');
            taskElement.className = 'task-item';
            taskElement.innerHTML = `
                <span>${task.name}</span>
                <div class="task-actions">
                    <button class="delete-btn" data-id="${task.id}">Delete</button>
                </div>
            `;
            taskList.appendChild(taskElement);
        });

        // Добавляем обработчики для кнопок удаления
        document.querySelectorAll('.delete-btn').forEach(button => {
            button.addEventListener('click', function() {
                const taskId = parseInt(this.getAttribute('data-id'));
                if (confirm('Are you sure you want to delete this task?')) {
                    deleteTask(taskId);
                }
            });
        });
    }

    function showMessage(message, type = 'info') {
        const msgElement = document.createElement('div');
        msgElement.className = `message ${type}`;
        msgElement.textContent = message;
        
        // Очищаем предыдущие сообщения
        const oldMessages = taskList.querySelectorAll('.message');
        oldMessages.forEach(msg => msg.remove());
        
        taskList.prepend(msgElement);
        
        // Автоматическое скрытие через 3 секунды (кроме ошибок авторизации)
        if (type !== 'error') {
            setTimeout(() => {
                msgElement.remove();
            }, 3000);
        }
    }
});