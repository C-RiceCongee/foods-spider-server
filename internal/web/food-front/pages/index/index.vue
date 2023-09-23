<template>
	<view>
		<u-search ref="search" placeholder="今天想食点撒子喔～" shape="square" :showAction="false"
			v-model="searchParams.foodName" @search="handlerSearch"></u-search>
		<view v-show="searchList.length===0" class="searchHistory">
			<u-text text="搜索历史"></u-text>
			<view v-if="searchHistory.length>0">
				<view class="searchItem" v-for="(item,index) in searchHistory" key="index">
					<u-text :text="item" type="info"></u-text>
				</view>
			</view>
			<u-text v-else text="暂无搜索记录" type="info"></u-text>
		</view>
		<view v-show="searchList.length>0">
			<u-list @scrolltolower="scrolltolower" :height="searchListUI.height">
				<u-list-item v-for="(item, index) in searchList" :key="index">
					<u-cell :title="item.title" :label="item.desc" @click="goToDetails(item)">
						<span slot="right-icon">{{item.brand}}</span>
					</u-cell>
				</u-list-item>
			</u-list>
		</view>
	</view>
</template>

<script>
	import {
		searchFoods
	} from "@/request/api.js"
	import {
		h
	} from "vue"
	export default {
		data() {
			return {
				searchParams: {
					foodName: "",
					pageNo: 0,
				},
				searchList: [],
				searchListUI: {
					height: 0
				},
				loadMore: "loadmore"
			}
		},
		computed: {
			searchHistory() {
				return this.$store.state.searchHistory
			}
		},
		onLoad() {
			uni.getSystemInfo().then(res => {
				this.searchListUI.height = res.windowHeight - this.$refs.search.height
			})
		},
		methods: {
			goToDetails(item) {
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
				console.log(value)
				value !== "" && this.$store.commit("setSearchHistory", value)
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
	.searchHistory {
		padding: 10px;
		box-sizing: border-box;

		.searchItem {
			margin: 10px 0;
		}
	}
</style>