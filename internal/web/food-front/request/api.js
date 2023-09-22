import {
	fetchGET
} from "./request"

export const searchFoods = params => fetchGET(`/api/v1/foods/search`, params)