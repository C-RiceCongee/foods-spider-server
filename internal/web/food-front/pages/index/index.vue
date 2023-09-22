<template>
	<view>
		<u-search ref="search" placeholder="今天想食点撒子喔～" shape="square" :showAction="false"
			v-model="searchParams.foodName" @search="handlerSearch"></u-search>
		<div class=" list">
			<u-list @scrolltolower="scrolltolower" :height="searchListUI.height">
				<u-list-item v-for="(item, index) in searchList" :key="index">
					<u-cell :title="item.title" :label="item.desc"  @click="goToDetails(item)">
						<span slot="right-icon">{{item.brand}}</span>
					</u-cell>
				</u-list-item>
			</u-list>
		</div>
	</view>
</template>

<script>
	import {
		searchFoods
	} from "@/request/api.js"
	export default {
		data() {
			return {
				searchParams: {
					foodName: "",
					pageNo: 0,
				},
				searchList: [{
					"title": "牛油火锅底料",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E7%89%9B%E6%B2%B9%E7%81%AB%E9%94%85%E5%BA%95%E6%96%99/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 749千卡 | 脂肪: 80.60克 | 碳水物: 6.10克 | 蛋白质: 2.80克"
				}, {
					"title": "龙口粉丝",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E9%BE%99%E5%8F%A3%E7%B2%89%E4%B8%9D/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 359千卡 | 脂肪: 0.00克 | 碳水物: 88.40克 | 蛋白质: 0.00克"
				}, {
					"title": "虾滑",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E8%99%BE%E6%BB%91/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 77千卡 | 脂肪: 0.00克 | 碳水物: 2.40克 | 蛋白质: 15.70克"
				}, {
					"title": "番茄火锅底料",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E7%95%AA%E8%8C%84%E7%81%AB%E9%94%85%E5%BA%95%E6%96%99/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 220千卡 | 脂肪: 12.70克 | 碳水物: 22.60克 | 蛋白质: 3.90克"
				}, {
					"title": "湿粉条",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E6%B9%BF%E7%B2%89%E6%9D%A1/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 193千卡 | 脂肪: 0.00克 | 碳水物: 47.50克 | 蛋白质: 0.00克"
				}, {
					"title": "原味火锅蘸料",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E5%8E%9F%E5%91%B3%E7%81%AB%E9%94%85%E8%98%B8%E6%96%99/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 232千卡 | 脂肪: 16.20克 | 碳水物: 12.50克 | 蛋白质: 9.40克"
				}, {
					"title": "牛肉包",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E7%89%9B%E8%82%89%E5%8C%85/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 320千卡 | 脂肪: 24.80克 | 碳水物: 1.00克 | 蛋白质: 23.70克"
				}, {
					"title": "虎牙脆",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E8%99%8E%E7%89%99%E8%84%86/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 530千卡 | 脂肪: 33.90克 | 碳水物: 50.50克 | 蛋白质: 6.10克"
				}, {
					"title": "黄金玉米花",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E9%BB%84%E9%87%91%E7%8E%89%E7%B1%B3%E8%8A%B1/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 533千卡 | 脂肪: 33.50克 | 碳水物: 52.40克 | 蛋白质: 6.00克"
				}, {
					"title": "自热米饭",
					"brand": "(海底捞)",
					"link": "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/%E6%B5%B7%E5%BA%95%E6%8D%9E/%E8%87%AA%E7%83%AD%E7%B1%B3%E9%A5%AD/100%E5%85%8B",
					"desc": "每100克 - 卡路里: 235千卡 | 脂肪: 7.60克 | 碳水物: 36.80克 | 蛋白质: 4.50克"
				}],
				searchListUI: {
					height: 0
				},
				loadMore: "loadmore"
			}
		},
		onLoad() {
			uni.getSystemInfo().then(res => {
				this.searchListUI.height = res.windowHeight - this.$refs.search.height
			})
		},
		methods: {
			goToDetails(item) {
				console.log(item)
				uni.navigateTo({
					url: `/pages/details/details?link=${item.link}`,
				})
			},
			scrolltolower() {
				if (this.loadMore === "loading") return
				this.loadMore = true
				this.searchParams.pageNo += 1
				this.fetchSearchByParams()
			},
			async handlerSearch(value) {
				this.searchParams.pageNo = 0
				this.searchList = []
				this.fetchSearchByParams()
			},
			async fetchSearchByParams() {
				this.loadMore = "loading"
				const res = await searchFoods({
					foodName: this.searchParams.foodName || "",
					pageNo: this.searchParams.pageNo
				})
				this.searchList.push(...res.result)
				this.loadMore = "loadmore"
			}
		}
	}
</script>

<style lang="scss" scoped>
	/deep/ .u-search {}
</style>