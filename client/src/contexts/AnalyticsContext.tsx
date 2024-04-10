import { createContext, useMemo } from "react";
import * as amplitude from "@amplitude/analytics-browser";

import { useCurrentUser } from "../hooks/useCurrentUser";

type AnalyticsContextType = {
  trackEvent: (
    name: string,
    attr?: Record<string, string | boolean | number>,
  ) => void;
};

const defaultAnalyticsContext: AnalyticsContextType = {
  trackEvent: () => {},
};

export const AnalyticsContext = createContext<AnalyticsContextType>(
  defaultAnalyticsContext,
);

type Props = {
  children: React.ReactNode;
};

export const AnalyticsProvider = ({ children }: Props) => {
  const { user } = useCurrentUser();

  const contextValue = useMemo(() => {
    if (import.meta.env.PROD) {
      amplitude.init(import.meta.env.VITE_AMPLITUDE_API_KEY, {
        defaultTracking: true,
      });

      if (user && user.id && user.email) {
        amplitude.setUserId(user.id);
        const identifyEvent = new amplitude.Identify();
        identifyEvent.set("email", user.email);
        amplitude.identify(identifyEvent);
      }
    }
    return { trackEvent: amplitude.logEvent };
  }, [user]);

  return (
    <AnalyticsContext.Provider value={contextValue}>
      {children}
    </AnalyticsContext.Provider>
  );
};
