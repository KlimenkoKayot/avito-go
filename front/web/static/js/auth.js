document.addEventListener('DOMContentLoaded', function() {
    const authForm = document.getElementById('loginForm') || document.getElementById('registerForm');
    
    if (authForm) {
        authForm.addEventListener('submit', async function(e) {
            e.preventDefault();
            
            // Очищаем предыдущие ошибки
            clearErrors();
            
            // Получаем данные формы
            const formData = {
                username: this.querySelector('#username')?.value.trim() || 
                         this.querySelector('#email')?.value.trim(),
                password: this.querySelector('#password').value.trim()
            };
            
            // Базовая валидация
            if (!formData.username || !formData.password) {
                showError('Заполните все обязательные поля');
                return;
            }
            
            // Для регистрации - проверка подтверждения пароля
            if (this.id === 'registerForm') {
                const confirmPassword = this.querySelector('#confirmPassword').value.trim();
                if (formData.password !== confirmPassword) {
                    showError('Пароли не совпадают', 'confirmPassword');
                    return;
                }
            }
            
            // Показываем индикатор загрузки
            setLoadingState(true);
            
            try {
                // Определяем endpoint
                const endpoint = this.id === 'loginForm' ? 'http://127.0.0.1:8081/auth/login' : 'http://127.0.0.1:8081/auth/register';
                
                // Отправляем запрос с таймаутом 5 секунд
                const response = await fetchWithTimeout(endpoint, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Accept': 'application/json',
                        'Origin': 'http://127.0.0.1:8080'
                    },
                    body: JSON.stringify({
                        login: formData.username,
                        pass: formData.password
                    }),
                    credentials: 'include'
                }, 5000);
                
                // Обрабатываем ответ
                await handleResponse(response);
                
            } catch (error) {
                console.error('Ошибка:', error);
                if (error.name === 'AbortError') {
                    showError('Сервер не отвечает. Попробуйте позже.');
                } else {
                    showError('Произошла ошибка. Пожалуйста, попробуйте еще раз.');
                }
            } finally {
                // Скрываем индикатор загрузки
                setLoadingState(false);
            }
        });
    }
    
    // Функция для fetch с таймаутом
    async function fetchWithTimeout(url, options, timeout) {
        const controller = new AbortController();
        const timeoutId = setTimeout(() => controller.abort(), timeout);
        
        try {
            const response = await fetch(url, {
                ...options,
                signal: controller.signal
            });
            clearTimeout(timeoutId);
            return response;
        } catch (error) {
            clearTimeout(timeoutId);
            throw error;
        }
    }
    
    // Обработчик ответа сервера
    async function handleResponse(response) {
        const contentType = response.headers.get('content-type');
        let responseData;
        
        if (contentType && contentType.includes('application/json')) {
            responseData = await response.json();
        }
        
        switch (response.status) {
            case 200: // Успех
                window.location.href = responseData?.redirect || '/';
                break;
                
            case 401: // Неавторизован
                showError(responseData?.message || 'Неверные учетные данные');
                break;
                
            case 404: // Техработы/не найдено
                showError(responseData?.message || 'Сервис временно недоступен. Идут технические работы.');
                break;
                
            default: // Другие ошибки
                throw new Error(responseData?.message || `HTTP error! status: ${response.status}`);
        }
    }
    
    // Показать ошибку
    function showError(message, fieldId = null) {
        if (fieldId) {
            const field = document.getElementById(fieldId);
            if (field) {
                field.classList.add('is-invalid');
                const feedback = field.nextElementSibling || document.createElement('div');
                feedback.className = 'invalid-feedback';
                feedback.textContent = message;
                field.parentNode.insertBefore(feedback, field.nextSibling);
            }
        } else {
            const alertDiv = document.createElement('div');
            alertDiv.className = 'alert alert-danger mt-3';
            alertDiv.textContent = message;
            authForm.parentNode.insertBefore(alertDiv, authForm.nextSibling);
            
            // Автоудаление через 5 секунд
            setTimeout(() => {
                alertDiv.remove();
            }, 5000);
        }
    }
    
    // Очистить ошибки
    function clearErrors() {
        document.querySelectorAll('.is-invalid').forEach(el => {
            el.classList.remove('is-invalid');
        });
        
        document.querySelectorAll('.invalid-feedback').forEach(el => {
            el.remove();
        });
        
        document.querySelectorAll('.alert').forEach(el => {
            el.remove();
        });
    }
    
    // Управление состоянием загрузки
    function setLoadingState(isLoading) {
        const submitButton = authForm.querySelector('button[type="submit"]');
        const submitText = submitButton.querySelector('.submit-text');
        const spinner = submitButton.querySelector('.spinner-border');
        
        submitButton.disabled = isLoading;
        
        if (isLoading) {
            submitText.style.opacity = '0';
            spinner.classList.remove('d-none');
        } else {
            submitText.style.opacity = '1';
            spinner.classList.add('d-none');
        }
    }
});