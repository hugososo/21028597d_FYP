import { ExoticComponent, lazy } from "react";
import { PAGE_TYPE } from "features/auth";
import Publish from "pages/Publish";
import Explore from "pages/Explore";
import Rewards from "pages/Rewards";
import BookDetail from "pages/BookDetail";
import Governance from "pages/Governance";
import ProposalDetail from "pages/ProposalDetail";
import { SETTING_SECTION } from "features/user/components/ProfileSideBar/ProfileSideBar";
import MyItems from "pages/Profile/MyItems";

const LandingPage = lazy(() => import("pages/LandingPage"));
const Empty = lazy(() => import("pages/Empty"));

type PageComponent = ExoticComponent | React.FC | React.Component;

type Page = {
  key: string;
  path: string;
  component: PageComponent;
  type: PAGE_TYPE;
};

export const mainRouteComponentsMapper: Page[] = [
  {
    key: "/",
    path: "/",
    component: LandingPage,
    type: PAGE_TYPE.PUBLIC,
  },
  {
    key: "publish",
    path: "/publish",
    component: Publish,
    type: PAGE_TYPE.PUBLIC,
  },
  {
    key: "explore",
    path: "/explore",
    component: Explore,
    type: PAGE_TYPE.PUBLIC,
  },
  {
    key: "explore",
    path: "/explore/:cid",
    component: BookDetail,
    type: PAGE_TYPE.PUBLIC,
  },
  {
    key: "reward",
    path: "/reward",
    component: Rewards,
    type: PAGE_TYPE.PUBLIC,
  },
  {
    key: "governance",
    path: "/governance",
    component: Governance,
    type: PAGE_TYPE.PUBLIC,
  },
  {
    key: "proposal",
    path: "/governance/proposal/:id",
    component: ProposalDetail,
    type: PAGE_TYPE.PUBLIC,
  },
];

export const profileRouteComponentsMapper = [
  {
    key: SETTING_SECTION.BOOKS,
    route: `/profile/:id/${SETTING_SECTION.BOOKS}`,
    component: MyItems,
  },
  {
    key: SETTING_SECTION.PUBLISHES,
    route: `/profile/:id/${SETTING_SECTION.PUBLISHES}`,
    component: MyItems,
  },
  {
    key: SETTING_SECTION.ACTIVITIES,
    route: `/profile/:id/${SETTING_SECTION.ACTIVITIES}`,
    component: Empty,
  },
];
