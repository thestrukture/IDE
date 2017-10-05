 //
//  FlowTissue.m
//  GoTetst2
//
//  Created by OrendaCapital on 12/29/15.
//  Copyright © 2015 Cheikh Seck LLC. All rights reserved.
//

#import "FlowTissue.h"
#import "FlowBluetooth.h"
#import "FlowAccellerometer.h"



@implementation FlowTissue



- (void) trackMotion {
    [[UIAccelerometer sharedAccelerometer] setDelegate:[FlowThreadManager instance]];
    NSLog(@"Watching movements");

}

+ (void) handleRequest:(NSString *) endpoint {
	GoMymobileLoadUrl(endpoint, nil, @"GET",[FlowThreadManager tissue]);
}

- (void) stopMotion {
    [[UIAccelerometer sharedAccelerometer] setDelegate:nil];
}


- (void) notify:(NSString *)title message:(NSString *)message {
    UILocalNotification* localNotification = [[UILocalNotification alloc] init];
    localNotification.fireDate = [NSDate dateWithTimeIntervalSinceNow:0];
    localNotification.alertBody = message;
    localNotification.alertTitle = title;
    localNotification.timeZone = [NSTimeZone defaultTimeZone];
    [[UIApplication sharedApplication] scheduleLocalNotification:localNotification];
}

/*
    Flow Tissue Core Comm between Go and native langs to reach hardware specs
    Sound, touch scan, app links, GPS and files...
*/

- (int) device {
    return 0;
}

- (void) createPictureNamed:(NSString *)name {
    //take picture and save to specified name....
       dispatch_async(dispatch_get_main_queue(), ^{
           [FlowThreadManager takePicture:name];
       });
    
}

//sound
- (void) play:(NSString *)path {
    
    NSError *error = nil;
    FlowThreadManager *stream = [FlowThreadManager instance];
    NSData *fileData = [NSData dataWithContentsOfFile:[[FlowTissue applicationDocumentsDirectory] stringByAppendingString:path] ];
    
    if (stream.audioPlayer != nil) {
        if (stream.isPlaying){
            [stream.audioPlayer stop];
        }
    }
    
    stream.audioPlayer = [[AVAudioPlayer alloc] initWithData:fileData error:&error];
    
    [stream.audioPlayer prepareToPlay];
    [stream.audioPlayer play];
    if (error == nil)
    stream.isPlaying = YES;
    else stream.isPlaying = NO;
}

- (void) playFromWebRoot:(NSString *)path {
    NSError *error = nil;
    FlowThreadManager *stream = [FlowThreadManager instance];
    NSData *fileData = GoMymobileLoadUrl(path, nil, @"GET", nil);
    
    if (stream.audioPlayer != nil) {
        if (stream.isPlaying){
            [stream.audioPlayer stop];
        }
    }
    
    stream.audioPlayer = [[AVAudioPlayer alloc] initWithData:fileData error:&error];
    
    [stream.audioPlayer prepareToPlay];
    [stream.audioPlayer play];
    
    if (error == nil)
    stream.isPlaying = YES;
    else stream.isPlaying = NO;
    
}

- (void) setVolume:(int)power {
    FlowThreadManager *stream = [FlowThreadManager instance];
    [stream.audioPlayer setVolume: (float) (power/100) ];
}

- (int) getVolume {
    FlowThreadManager *stream = [FlowThreadManager instance];
    //[stream.audioPlayer setVolume: (float) (power/100) ];
    return 100*stream.audioPlayer.volume;
}

- (void) stop {
    FlowThreadManager *stream = [FlowThreadManager instance];
    stream.isPlaying = NO;
    [stream.audioPlayer stop];
}

- (BOOL) isPlaying {
    FlowThreadManager *stream = [FlowThreadManager instance];
    return stream.isPlaying;
}


//Applinks
- (void) openAppLink:(NSString *)url {
        //process applinkios
    dispatch_async(dispatch_get_main_queue(), ^{
    UIApplication *ourApplication = [UIApplication sharedApplication];
    NSString *URLEncodedText = [url stringByAddingPercentEscapesUsingEncoding:NSUTF8StringEncoding];
    NSString *ourPath =URLEncodedText;
    NSURL *ourURL = [NSURL URLWithString:ourPath];
    if ([ourApplication canOpenURL:ourURL]) {
        [ourApplication openURL:ourURL];
    }
    });
}

//GPS
- (void) requestLocation {
    //[[FlowThreadManager getGPS] requestWhenInUseAuthorization];
    //[[FlowThreadManager getGPS] requestLocation];
}

- (void) showLoad {
    dispatch_async(dispatch_get_main_queue(), ^{
    [FlowThreadManager loadScreen:YES usingMessage:@""];
    });
}

- (void) hideLoad {
    [FlowThreadManager loadScreen:NO usingMessage:@""];
}

- (void) runJS:(NSString *)line {
    dispatch_async(dispatch_get_main_queue(), ^{
    [FlowThreadManager runJS:line];
    });
}




//files
- (NSString *) absolutePath:(NSString *)file {
    return [[FlowTissue applicationDocumentsDirectory] stringByAppendingString:file];
}

- (BOOL) download:(NSString *)url target:(NSString *)target {
    
    //NSString *stringURL = @"http://www.somewhere.com/thefile.png";
    NSURL  *urll = [NSURL URLWithString:url];
    NSData *urlData = [NSData dataWithContentsOfURL:urll];
    if ( urlData )
    {
        NSString  *filePath = [self absolutePath:target];
        [urlData writeToFile:filePath atomically:YES];
        return YES;
    }
    
    return NO;
}

- (void) downloadLarge:(NSString *)url target:(NSString *)target {
    
    dispatch_async(dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_DEFAULT, 0), ^{

    //NSString *stringURL = @"http://www.somewhere.com/thefile.png";
    NSURL  *urll = [NSURL URLWithString:url];
    NSData *urlData = [NSData dataWithContentsOfURL:urll];
    if ( urlData )
    {
        NSString  *filePath = [self absolutePath:target];
        dispatch_async(dispatch_get_main_queue(), ^{
        [urlData writeToFile:filePath atomically:YES];
        });
       
    }
        
    });
   
}

- (NSString *) base64String:(NSString *)target {
    return [[self getBytes:target] base64EncodedStringWithOptions:0];
}

- (NSData *) getBytes:(NSString *)target {
    return [NSData dataWithContentsOfFile:[self absolutePath:target]];
}

- (NSData *) getBytesFromUrl:(NSString *)target {
    return [NSData dataWithContentsOfURL:[NSURL URLWithString:[self absolutePath:target]]];
}


- (BOOL) deleteDirectory:(NSString *)path {
    return [[NSFileManager defaultManager] removeItemAtPath:[self absolutePath:path] error:nil];

}

- (BOOL) deleteFile:(NSString *)path {
    return [self deleteDirectory:path];
}






+ (NSString *) applicationDocumentsDirectory
{
    NSArray *paths = NSSearchPathForDirectoriesInDomains(NSDocumentDirectory, NSUserDomainMask, YES);
    NSString *basePath = paths.firstObject;
    return basePath;
}


- (double) width {
    CGFloat width = [UIScreen mainScreen].bounds.size.width;
    return (double) width;
}

- (double) height {
    CGFloat height = [UIScreen mainScreen].bounds.size.height;
    return (double) height;
}


- (void)pushView:(NSString *)url {
        dispatch_async(dispatch_get_main_queue(), ^{
          [FlowThreadManager pulseView:url];
        });
    
    NSLog(@"Openning view %@", url);
}

- (void) dismissView {
    dispatch_async(dispatch_get_main_queue(), ^{
    UINavigationController *navcontroller = (UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController;
    // Replace the current view controller
    NSMutableArray *viewControllers = [NSMutableArray arrayWithArray:[navcontroller viewControllers]];
    
    [viewControllers removeLastObject];
    
    [navcontroller setViewControllers:viewControllers animated:YES];
    });
}

- (void) dismissViewatInt:(int)index {
     dispatch_async(dispatch_get_main_queue(), ^{
    UINavigationController *navcontroller = (UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController;
    // Replace the current view controller
    NSMutableArray *viewControllers = [NSMutableArray arrayWithArray:[navcontroller viewControllers]];
    
    [viewControllers removeObjectAtIndex:index];
    
    [navcontroller setViewControllers:viewControllers animated:YES];
         
    });
}


@end
