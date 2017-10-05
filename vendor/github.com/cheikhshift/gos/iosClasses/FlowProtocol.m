//
					//  FlowProtocol.m
					//  FlowCode
					//
					//  Created by Cheikh Seck on 4/2/15.
					//  Copyright (c) 2015 Gopher Sauce LLC. All rights reserved.
					//

					#import "FlowProtocol.h"
					#import "FlowTissue.h"
					#import "Mymobile/Mymobile.h"

					@implementation FlowProtocol



					+ (BOOL)canInitWithRequest:(NSURLRequest*)theRequest
					{
					    if ([theRequest.URL.host caseInsensitiveCompare:@"localhost"] == NSOrderedSame) {
					        return YES;
					    }
					    return NO;
					}

					+ (NSURLRequest*)canonicalRequestForRequest:(NSURLRequest*)theRequest
					{
					    return theRequest;
					}

					- (void)startLoading
					{
					  
					    NSString *process = [self.request.URL.absoluteString stringByReplacingOccurrencesOfString:@"http://localhost" withString:@""];
					    //check here
					    NSString *GetString;
					   //NSLog(@"%@", self.request.HTTPBody );
					    if([process rangeOfString:@"?"].location != NSNotFound){
					        if([process componentsSeparatedByString:@"?"].count > 1 )
					        GetString = [[process componentsSeparatedByString:@"?"] objectAtIndex:1];
					        process = [[process componentsSeparatedByString:@"?"] objectAtIndex:0];
					    }


                        if([self.request HTTPBody] != nil && [self.request.HTTPBody length] > 0){
                            GetString = [GetString stringByAppendingString:@"&"];
                            GetString = [GetString stringByAppendingString:[NSString stringWithUTF8String:[self.request.HTTPBody bytes] ]];
                        }
					    
					    CFStringRef fileExtension = (__bridge CFStringRef)[process pathExtension];
					    CFStringRef UTI = UTTypeCreatePreferredIdentifierForTag(kUTTagClassFilenameExtension, fileExtension, NULL);
					    CFStringRef MIMEType = UTTypeCopyPreferredTagWithClass(UTI, kUTTagClassMIMEType);
					    CFRelease(UTI);
					    NSString *MIMETypeString = (__bridge_transfer NSString *)MIMEType;
					    NSURLResponse *response = [[NSURLResponse alloc] initWithURL:[self.request URL]
					                                                        MIMEType:MIMETypeString
					                                           expectedContentLength:-1
					                                                textEncodingName:nil];
					    
					      dispatch_async(dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_DEFAULT, 0), ^{
					          
					    //NSLog(@"%@", self.request.HTTPBody );
					   
					          
					  
					    [[self client] URLProtocol:self didReceiveResponse:response cacheStoragePolicy:NSURLCacheStorageNotAllowed];
					   
					    [[self client] URLProtocol:self didLoadData:GoMymobileLoadUrl(process, [self parseParams:GetString], self.request.HTTPMethod,[FlowThreadManager tissue])];
					    [[self client] URLProtocolDidFinishLoading:self];
					      });
					   
					}

					- (NSData *) parseParams: (NSString *) input {
					    if(![input isEqualToString:@""]){
					    NSArray *pieces = [input componentsSeparatedByString:@"&"];
					    NSDictionary *payload = [NSMutableDictionary new];
					    
					    
					    
					    for (int i = 0; i < pieces.count; i++) {
					        NSString * param  = [pieces objectAtIndex:i];
					        if(![param isEqualToString:@""]){
					         
					            NSArray *keyset = [param componentsSeparatedByString:@"="];
					            [payload setValue:[self urlDecode:[keyset objectAtIndex:1] ] forKey:[self urlDecode:[keyset objectAtIndex:0]] ];
					            
					        }
					    }
					    NSError *error;
					    NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload
					                                                       options:NSJSONWritingPrettyPrinted // Pass 0 if you don't care about the readability of the generated string
					                                                         error:&error];
					    
					    if (! jsonData) {
					        NSLog(@"Got an error: %@", error);
					        return nil;
					    } else {
					        NSString *jsonString = [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];
					        return [jsonString dataUsingEncoding:NSUTF8StringEncoding];
					    }
					    }
					    return nil;
					    
					}

					- (NSString *) urlDecode :(NSString *) input {
					    return [[input stringByReplacingOccurrencesOfString:@"+" withString:@" "]
					            stringByReplacingPercentEscapesUsingEncoding:NSUTF8StringEncoding];
					}
	

					- (void) stopLoading {
					    
					}

					@end
