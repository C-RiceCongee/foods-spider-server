import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
const SEARCH_HISTORY = "searchHistory"
const his = uni.getStorageSync(SEARCH_HISTORY)

const store = new Vuex.Store({
	state: {
		searchHistory: his ? JSON.parse(his) : [],
	},
	mutations: {
		setSearchHistory(state, params) {
			state.searchHistory.push(params)
			uni.setStorageSync(SEARCH_HISTORY, JSON.stringify(state.searchHistory))
		}
	}
})
export default store