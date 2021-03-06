package app.bigants.rn;

import com.facebook.react.ReactActivity;
import com.facebook.react.ReactActivityDelegate; // ADD: (react-navigation) https://reactnavigation.org/docs/en/getting-started.html
import com.facebook.react.ReactRootView; // ADD: (react-navigation) https://reactnavigation.org/docs/en/getting-started.html
import com.swmansion.gesturehandler.react.RNGestureHandlerEnabledRootView; // ADD: (react-navigation) https://reactnavigation.org/docs/en/getting-started.html

public class MainActivity extends ReactActivity {

    /**
     * Returns the name of the main component registered from JavaScript.
     * This is used to schedule rendering of the component.
     */
    @Override
    protected String getMainComponentName() {
        return "bigants";
    }

    // ADD: (react-navigation) <--
    @Override
    protected ReactActivityDelegate createReactActivityDelegate() {
        return new ReactActivityDelegate(this, getMainComponentName()) {
        @Override
        protected ReactRootView createRootView() {
            return new RNGestureHandlerEnabledRootView(MainActivity.this);
        }
    };
    }   
    // -->  https://reactnavigation.org/docs/en/getting-started.html
}
