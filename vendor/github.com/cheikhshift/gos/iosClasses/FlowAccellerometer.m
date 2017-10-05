//
//  FlowAccellerometer.m
//  GoTetst2
//
//  Created by OrendaCapital on 1/9/16.
//  Copyright © 2016 Cheikh Seck LLC. All rights reserved.
//

#import "FlowAccellerometer.h"
#import "FlowThreadManager.h"

@implementation FlowAccellerometer

-(void) stop {
    //dc!!!!
    [[UIAccelerometer sharedAccelerometer] setDelegate:nil];
}

- (void) start {
    [[UIAccelerometer sharedAccelerometer] setDelegate:[FlowThreadManager instance]];
    NSLog(@"Watching movements");
}

- (void) calibrate {
    
}

@end
