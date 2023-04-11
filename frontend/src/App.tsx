import Loading from "components/UI/Loading";
import { ExoticComponent, Suspense } from "react";
import { Route, Routes } from "react-router-dom";
import { mainRouteComponentsMapper, profileRouteComponentsMapper } from "routes";
import { PAGE_TYPE } from "features/auth";
import Web3ClientProvider from "providers/Web3ClientProvider";
import NotificationProvider from "features/notificaiton/providers/NotificationProvider";
import { Redirect, PageNavigator } from "features/navigation/components";
import { ProfileSideBar } from "features/user/components";
import Main from "pages/Main";
import ContractInstanceProvider from "providers/ContractInstanceProvider";
import Center from "components/UI/Center";
import { HttpClientProvider } from "providers/HttpClientProvider";
import Profile from "pages/Profile";

const App = () => {
  return (
    <NotificationProvider>
      <Web3ClientProvider>
        <ContractInstanceProvider>
          <HttpClientProvider>
            <Suspense
              fallback={
                <Center>
                  <Loading />
                </Center>
              }
            >
              <Routes>
                <Route
                  element={
                    <PageNavigator pageType={PAGE_TYPE.PUBLIC}>
                      <Main />
                    </PageNavigator>
                  }
                >
                  {mainRouteComponentsMapper.map(({ key, path, component }) => {
                    const Component = component as ExoticComponent;
                    return <Route key={key} path={path} element={<Component />} />;
                  })}
                  <Route
                    element={
                      <PageNavigator pageType={PAGE_TYPE.PRIVATE}>
                        <Profile />
                      </PageNavigator>
                    }
                  >
                    {profileRouteComponentsMapper.map(({ key, route, component }) => {
                      const Component = component as ExoticComponent;
                      return <Route key={key} path={route} element={<Component />} />;
                    })}
                  </Route>
                </Route>
                <Route path="*" element={<Redirect />} />
              </Routes>
            </Suspense>
          </HttpClientProvider>
        </ContractInstanceProvider>
      </Web3ClientProvider>
    </NotificationProvider>
  );
};

export default App;
