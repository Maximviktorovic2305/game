import axios from 'axios'

// Базовый URL для API
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

// Клиент для выполнения HTTP-запросов
const apiClient = axios.create({
	baseURL: `${API_BASE_URL}/api`,
	headers: {
		'Content-Type': 'application/json',
	},
})

export default apiClient
