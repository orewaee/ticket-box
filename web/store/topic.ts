import {defineStore} from "pinia";
import axios from "axios";
import type {Topic, TopicState} from "./types";

const baseUrl = "http://localhost:8083";

export const useTopicState = defineStore("topic", {
    state: (): TopicState => ({
        topic: null,
        error: false,
        loading: true
    }),
    getters: {
        id: state => state.topic?.id,
        guildId: state => state.topic?.guild_id,
        emoji: state => state.topic?.emoji,
        name: state => state.topic?.name,
        description: state => state.topic?.description,
    },
    actions: {
        async getTopicById(topicId: string) {
            this.loading = true;

            try {
                const {data} = await axios.get<Topic>(baseUrl + "/topic/" + topicId, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                });
                this.topic = data;
                this.error = false;
            } catch (error) {
                this.error = true;
            } finally {
                this.loading = false;
            }
        },
        async reset() {
            this.topic = null
            this.error = false
            this.loading = true
        }
    }
});