import {
	fetchGET
} from "./request"

export const searchFoods = params => fetchGET(`/api/v1/foods/search`, params)
export const linkDetails = params => fetchGET(`/api/v1/foods/linkDetails`, params)