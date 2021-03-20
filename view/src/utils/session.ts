import Vue from 'vue';

export default class Session extends Vue {
    public remove(sessionKey: string) {
        if (this.$cookies.isKey(sessionKey)) {
            this.$cookies.remove(sessionKey);
        }
    }

    public exist(sessionKey: string) {
        return this.$cookies.isKey(sessionKey);
    }
}
