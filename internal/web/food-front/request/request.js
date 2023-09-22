export const baseURL = 'http://localhost:8888'

export const fetchGET = (url, params) => request({
	url,
	data: params,
})

export default function request(options) {
	return new Promise((resolve, reject) => {
		uni.request({
			url: baseURL + options.url,
			method: options.method || 'GET',
			data: options.data || {},
			header: options.header || {
				Authorization: "Bearer " + uni.getStorageSync("token")
			},
			success: res => {
				if (res.statusCode === 500) {
					uni.showToast({
						title: "服务器异常",
						icon: "error"
					})
					return
				}
				// 成功回调
				const {
					code,
					message
				} = res.data
				if (code === 1000) {
					resolve(res.data)
				} else if (code === 1007) {
					store.commit('UPDATE_LOGIN_STATUS', false)
					uni.removeStorageSync("token")
				} else {
					uni.showToast({
						title: message,
						icon: "error"
					})
					reject(message)
				}
			},
			fail: err => {
				uni.showToast({
					title: "服务器异常",
					icon: "error"
				})
				console.log("err")
				// 失败回调
				reject(err)
			}
		})
	})
}