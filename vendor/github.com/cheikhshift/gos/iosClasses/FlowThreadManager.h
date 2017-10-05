//
//  FlowThreadManager.h
//  FlowCode
//
//  Created by Cheikh Seck on 4/1/15.
//  Copyright (c) 2015 Gopher Sauce LLC. All rights reserved.
//

#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>
#import <AVFoundation/AVFoundation.h>
#import "ViewController.h"
#import <CoreLocation/CoreLocation.h>
#import "FlowTissue.h"
#import "DejalActivityView.h"

@interface FlowThreadManager : NSObject <UIWebViewDelegate,CLLocationManagerDelegate,UIAccelerometerDelegate,UIImagePickerControllerDelegate>


typedef void(^Completion)(void);
typedef void(^CodeProcessCompletion)(NSString *result);

@property (nonatomic) BOOL isPlaying;
@property (nonatomic, strong) AVAudioPlayer *audioPlayer;
@property (nonatomic) NSString *tempstring;
@property (strong, nonatomic) CLLocationManager *locationManager;

+ (void) getGPS;
+ (void) runJS:(NSString *) function;
+ (FlowThreadManager *) instance;
+ (id) tissue;
+ (void) loadScreen:(BOOL) switc usingMessage: (NSString *) message;
+ (UIWebView *) currentFlow;
+ (NSString *) flowjs:(NSString *)function withData:(NSArray *) args;
+ (void) createFlowLayer;
+ (void) process: (NSString *) mvc completion:(CodeProcessCompletion) finished;
+ (id) getobject:(NSString *) name;
+ (void) takePicture:(NSString *) name;
+ (BOOL) saveobject:(id)object withName:(NSString *) key;
+ (void) webviewCompletion:(Completion) finished;
+ (void) userDidCancelPayment;
+ (void) pulseView : (NSString *) url;
+ (void) pinLogin;

@end
