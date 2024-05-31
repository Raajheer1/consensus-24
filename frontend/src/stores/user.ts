import { defineStore } from "pinia";
import { API } from "@/utils/api";
import { Account } from "@/types";
import { notify } from "notiwind";

import useLoanStore from "@/stores/loan";

interface UserState {
  user: Account | null;
  public_key: string;
  secret_key: string;
  profile_pic: string;
  error: string | null;
  fetching: boolean;
  hasFetched: boolean;
  loading: Promise<void> | null;
}

const useUserStore = defineStore({
  id: "user",
  state: () =>
    ({
      user: null,
      public_key: "FDSJFSDKLDFSJFDSLFKLJS",
      secret_key: "",
      error: null,
      fetching: false,
      hasFetched: false,
    }) as UserState,
  getters: {
    isLoggedIn: (state) => !!state.user,
    fullName: (state) => {
      if (!state.user) return "";
      return `${state.user.first_name} ${state.user.last_name}`;
    },
  },
  actions: {
    async fetchUser(): Promise<void> {
      this.hasFetched = true;
      this.fetching = true;
      try {
        const { data } = await API.get(`/v1/account/${this.public_key}`);
        this.user = data;
        notify(
          {
            group: "br",
            title: "Login Successful",
            text: "Welcome back to MicroStars!",
          },
          4000
        );
        const loanStore = useLoanStore();
        await loanStore.fetchMyLoans(data.id.toString());
      } catch (e) {
        this.user = null;
      } finally {
        this.fetching = false;
      }
    },
    async fetchBorrower(id: number): Promise<Account | null> {
      this.fetching = true;
      try {
        const { data } = await API.get(`/v1/account/${id}`);
        return data;
      } catch (e) {
        console.log(e);
        return null;
      } finally {
        this.fetching = false;
      }
    },
    async logout(): Promise<void> {
      this.fetching = true;
      try {
        await API.get("/user/logout");
        this.user = null;
        this.hasFetched = false;
      } catch (e) {
        // TODO - throw error notification
        console.log(e);
      } finally {
        this.fetching = false;
      }
    },
  },
});

export default useUserStore;
