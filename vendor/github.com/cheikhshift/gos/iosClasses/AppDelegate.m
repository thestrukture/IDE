//
//  AppDelegate.m
//  Go Mobile Test
//
//  Created by OrendaCapital on 12/13/15.
//  Copyright © 2015 Cheikh Seck LLC. All rights reserved.
//

#import "AppDelegate.h"
#import "ViewController.h"
#import "FlowTissue.h"

@interface AppDelegate ()

@end

@implementation AppDelegate {
    UINavigationController *navController;
}

-(void)application:(UIApplication *)application performFetchWithCompletionHandler:(void (^)(UIBackgroundFetchResult))completionHandler
{
    
    //Tell the system that you ar done. || or no new data in that frame push etc
    NSLog(@"BG event");
    [FlowTissue handleRequest:@"/background"];
    completionHandler(UIBackgroundFetchResultNewData);
}

- (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(NSDictionary *)launchOptions {
    // Override point for customization after application launch.
    
    
    self.window = [[UIWindow alloc] initWithFrame:UIScreen.mainScreen.bounds];
    [application registerUserNotificationSettings:[UIUserNotificationSettings settingsForTypes:UIUserNotificationTypeAlert|UIUserNotificationTypeBadge|UIUserNotificationTypeSound categories:nil]];
    
    UIStoryboard *st = [UIStoryboard storyboardWithName:@"Main"
                                                 bundle:[NSBundle mainBundle]];
    [[UIApplication sharedApplication] setMinimumBackgroundFetchInterval:UIApplicationBackgroundFetchIntervalMinimum];
    
    ViewController *viewController =  [st instantiateViewControllerWithIdentifier:@"FlowView"];
    navController = [[UINavigationController alloc] initWithRootViewController:viewController];
    //viewController.webView.delegate = [FlowThreadManager instance];
    [navController setNavigationBarHidden:YES];
    
    
    self.window.rootViewController = navController;
    [self.window makeKeyAndVisible];
    [self.window addSubview:navController.view];
    

    
    
    return YES;
}



- (void)applicationWillResignActive:(UIApplication *)application {
    // Sent when the application is about to move from active to inactive state. This can occur for certain types of temporary interruptions (such as an incoming phone call or SMS message) or when the user quits the application and it begins the transition to the background state.
    // Use this method to pause ongoing tasks, disable timers, and throttle down OpenGL ES frame rates. Games should use this method to pause the game.
}

- (void)applicationDidEnterBackground:(UIApplication *)application {
    // Use this method to release shared resources, save user data, invalidate timers, and store enough application state information to restore your application to its current state in case it is terminated later.
    // If your application supports background execution, this method is called instead of applicationWillTerminate: when the user quits.
}

- (void)applicationWillEnterForeground:(UIApplication *)application {
    // Called as part of the transition from the background to the inactive state; here you can undo many of the changes made on entering the background.
}

- (void)applicationDidBecomeActive:(UIApplication *)application {
    // Restart any tasks that were paused (or not yet started) while the application was inactive. If the application was previously in the background, optionally refresh the user interface.
}

- (void)applicationWillTerminate:(UIApplication *)application {
    // Called when the application is about to terminate. Save data if appropriate. See also applicationDidEnterBackground:.
}

@end
