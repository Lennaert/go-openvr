// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package openvr

// OpenVR Constants
const (
	DriverNone                                  = uint(4294967295)
	MaxDriverDebugResponseSize                  = uint(32768)
	TrackedDeviceIndexHmd                       = uint(0)
	MaxTrackedDeviceCount                       = uint(16)
	TrackedDeviceIndexOther                     = uint(4294967294)
	TrackedDeviceIndexInvalid                   = uint(4294967295)
	InvalidPropertyContainer                    = uint(0)
	InvalidPropertyTag                          = uint(0)
	FloatPropertyTag                            = uint(1)
	Int32PropertyTag                            = uint(2)
	Uint64PropertyTag                           = uint(3)
	BoolPropertyTag                             = uint(4)
	StringPropertyTag                           = uint(5)
	HmdMatrix34PropertyTag                      = uint(20)
	HmdMatrix44PropertyTag                      = uint(21)
	HmdVector3PropertyTag                       = uint(22)
	HmdVector4PropertyTag                       = uint(23)
	HiddenAreaPropertyTag                       = uint(30)
	OpenVRInternalReservedStart                 = 1000
	OpenVRInternalReservedEnd                   = 10000
	MaxPropertyStringSize                       = uint(32768)
	ControllerStateAxisCount                    = uint(5)
	OverlayHandleInvalid                        = uint(0)
	ScreenshotHandleInvalid                     = uint(0)
	IVRSystemVersion                            = "IVRSystem_017"
	IVRExtendedDisplayVersion                   = "IVRExtendedDisplay_001"
	IVRTrackedCameraVersion                     = "IVRTrackedCamera_003"
	MaxApplicationKeyLength                     = uint(128)
	MimeTypeHomeApp                             = "vr/home"
	MimeTypeGameTheater                         = "vr/game_theater"
	IVRApplicationsVersion                      = "IVRApplications_006"
	IVRChaperoneVersion                         = "IVRChaperone_003"
	IVRChaperoneSetupVersion                    = "IVRChaperoneSetup_005"
	IVRCompositorVersion                        = "IVRCompositor_021"
	VROverlayMaxKeyLength                       = uint(128)
	VROverlayMaxNameLength                      = uint(128)
	MaxOverlayCount                             = uint(64)
	MaxOverlayIntersectionMaskPrimitivesCount   = uint(32)
	IVROverlayVersion                           = "IVROverlay_016"
	ControllerComponentGDC2015                  = "gdc2015"
	ControllerComponentBase                     = "base"
	ControllerComponentTip                      = "tip"
	ControllerComponentHandGrip                 = "handgrip"
	ControllerComponentStatus                   = "status"
	IVRRenderModelsVersion                      = "IVRRenderModels_005"
	NotificationTextMaxSize                     = uint(256)
	IVRNotificationsVersion                     = "IVRNotifications_002"
	MaxSettingsKeyLength                        = uint(128)
	IVRSettingsVersion                          = "IVRSettings_002"
	SteamVRSection                              = "steamvr"
	SteamVRRequireHmdString                     = "requireHmd"
	SteamVRForcedDriverKeyString                = "forcedDriver"
	SteamVRForcedHmdKeyString                   = "forcedHmd"
	SteamVRDisplayDebugBool                     = "displayDebug"
	SteamVRDebugProcessPipeString               = "debugProcessPipe"
	SteamVREnableDistortionBool                 = "enableDistortion"
	SteamVRDisplayDebugXInt32                   = "displayDebugX"
	SteamVRDisplayDebugYInt32                   = "displayDebugY"
	SteamVRSendSystemButtonToAllAppsBool        = "sendSystemButtonToAllApps"
	SteamVRLogLevelInt32                        = "loglevel"
	SteamVRIPDFloat                             = "ipd"
	SteamVRBackgroundString                     = "background"
	SteamVRBackgroundUseDomeProjectionBool      = "backgroundUseDomeProjection"
	SteamVRBackgroundCameraHeightFloat          = "backgroundCameraHeight"
	SteamVRBackgroundDomeRadiusFloat            = "backgroundDomeRadius"
	SteamVRGridColorString                      = "gridColor"
	SteamVRPlayAreaColorString                  = "playAreaColor"
	SteamVRShowStageBool                        = "showStage"
	SteamVRActivateMultipleDriversBool          = "activateMultipleDrivers"
	SteamVRDirectModeBool                       = "directMode"
	SteamVRDirectModeEdidVidInt32               = "directModeEdidVid"
	SteamVRDirectModeEdidPidInt32               = "directModeEdidPid"
	SteamVRUsingSpeakersBool                    = "usingSpeakers"
	SteamVRSpeakersForwardYawOffsetDegreesFloat = "speakersForwardYawOffsetDegrees"
	SteamVRBaseStationPowerManagementBool       = "basestationPowerManagement"
	SteamVRNeverKillProcessesBool               = "neverKillProcesses"
	SteamVRSupersampleScaleFloat                = "supersampleScale"
	SteamVRRenderTargetMultiplierFloat          = "renderTargetMultiplier"
	SteamVRAllowAsyncReprojectionBool           = "allowAsyncReprojection"
	SteamVRAllowReprojectionBool                = "allowInterleavedReprojection"
	SteamVRForceReprojectionBool                = "forceReprojection"
	SteamVRForceFadeOnBadTrackingBool           = "forceFadeOnBadTracking"
	SteamVRDefaultMirrorViewInt32               = "defaultMirrorView"
	SteamVRShowMirrorViewBool                   = "showMirrorView"
	SteamVRMirrorViewGeometryString             = "mirrorViewGeometry"
	SteamVRStartMonitorFromAppLaunch            = "startMonitorFromAppLaunch"
	SteamVRStartCompositorFromAppLaunchBool     = "startCompositorFromAppLaunch"
	SteamVRStartDashboardFromAppLaunchBool      = "startDashboardFromAppLaunch"
	SteamVRStartOverlayAppsFromDashboardBool    = "startOverlayAppsFromDashboard"
	SteamVREnableHomeApp                        = "enableHomeApp"
	SteamVRSetInitialDefaultHomeApp             = "setInitialDefaultHomeApp"
	SteamVRCycleBackgroundImageTimeSecInt32     = "CycleBackgroundImageTimeSec"
	SteamVRRetailDemoBool                       = "retailDemo"
	SteamVRIpdOffsetFloat                       = "ipdOffset"
	SteamVRAllowSupersampleFilteringBool        = "allowSupersampleFiltering"
	SteamVREnableLinuxVulkanAsyncBool           = "enableLinuxVulkanAsync"
	LighthouseSection                           = "driver_lighthouse"
	LighthouseDisableIMUBool                    = "disableimu"
	LighthouseUseDisambiguationString           = "usedisambiguation"
	LighthouseDisambiguationDebugInt32          = "disambiguationdebug"
	LighthousePrimaryBasestationInt32           = "primarybasestation"
	LighthouseDBHistoryBool                     = "dbhistory"
	NullSection                                 = "driver_null"
	NullEnableNullDriverBool                    = "enable"
	NullSerialNumberString                      = "serialNumber"
	NullModelNumberString                       = "modelNumber"
	NullWindowXInt32                            = "windowX"
	NullWindowYInt32                            = "windowY"
	NullWindowWidthInt32                        = "windowWidth"
	NullWindowHeightInt32                       = "windowHeight"
	NullRenderWidthInt32                        = "renderWidth"
	NullRenderHeightInt32                       = "renderHeight"
	NullSecondsFromVsyncToPhotonsFloat          = "secondsFromVsyncToPhotons"
	NullDisplayFrequencyFloat                   = "displayFrequency"
	UserInterfaceSection                        = "userinterface"
	UserInterfaceStatusAlwaysOnTopBool          = "StatusAlwaysOnTop"
	UserInterfaceMinimizeToTrayBool             = "MinimizeToTray"
	UserInterfaceScreenshotsBool                = "screenshots"
	UserInterfaceScreenshotTypeInt              = "screenshotType"
	NotificationsSection                        = "notifications"
	NotificationsDoNotDisturbBool               = "DoNotDisturb"
	KeyboardSection                             = "keyboard"
	KeyboardTutorialCompletions                 = "TutorialCompletions"
	KeyboardScaleX                              = "ScaleX"
	KeyboardScaleY                              = "ScaleY"
	KeyboardOffsetLeftX                         = "OffsetLeftX"
	KeyboardOffsetRightX                        = "OffsetRightX"
	KeyboardOffsetY                             = "OffsetY"
	KeyboardSmoothing                           = "Smoothing"
	PerfSection                                 = "perfcheck"
	PerfHeuristicActiveBool                     = "heuristicActive"
	PerfNotifyInHMDBool                         = "warnInHMD"
	PerfNotifyOnlyOnceBool                      = "warnOnlyOnce"
	PerfAllowTimingStoreBool                    = "allowTimingStore"
	PerfSaveTimingsOnExitBool                   = "saveTimingsOnExit"
	PerfTestDataFloat                           = "perfTestData"
	PerfLinuxGPUProfilingBool                   = "linuxGPUProfiling"
	CollisionBoundsSection                      = "collisionBounds"
	CollisionBoundsStyleInt32                   = "CollisionBoundsStyle"
	CollisionBoundsGroundPerimeterOnBool        = "CollisionBoundsGroundPerimeterOn"
	CollisionBoundsCenterMarkerOnBool           = "CollisionBoundsCenterMarkerOn"
	CollisionBoundsPlaySpaceOnBool              = "CollisionBoundsPlaySpaceOn"
	CollisionBoundsFadeDistanceFloat            = "CollisionBoundsFadeDistance"
	CollisionBoundsColorGammaRInt32             = "CollisionBoundsColorGammaR"
	CollisionBoundsColorGammaGInt32             = "CollisionBoundsColorGammaG"
	CollisionBoundsColorGammaBInt32             = "CollisionBoundsColorGammaB"
	CollisionBoundsColorGammaAInt32             = "CollisionBoundsColorGammaA"
	CameraSection                               = "camera"
	CameraEnableCameraBool                      = "enableCamera"
	CameraEnableCameraInDashboardBool           = "enableCameraInDashboard"
	CameraEnableCameraForCollisionBoundsBool    = "enableCameraForCollisionBounds"
	CameraEnableCameraForRoomViewBool           = "enableCameraForRoomView"
	CameraBoundsColorGammaRInt32                = "cameraBoundsColorGammaR"
	CameraBoundsColorGammaGInt32                = "cameraBoundsColorGammaG"
	CameraBoundsColorGammaBInt32                = "cameraBoundsColorGammaB"
	CameraBoundsColorGammaAInt32                = "cameraBoundsColorGammaA"
	CameraBoundsStrengthInt32                   = "cameraBoundsStrength"
	AudioSection                                = "audio"
	AudioOnPlaybackDeviceString                 = "onPlaybackDevice"
	AudioOnRecordDeviceString                   = "onRecordDevice"
	AudioOnPlaybackMirrorDeviceString           = "onPlaybackMirrorDevice"
	AudioOffPlaybackDeviceString                = "offPlaybackDevice"
	AudioOffRecordDeviceString                  = "offRecordDevice"
	AudioVIVEHDMIGain                           = "viveHDMIGain"
	PowerSection                                = "power"
	PowerPowerOffOnExitBool                     = "powerOffOnExit"
	PowerTurnOffScreensTimeoutFloat             = "turnOffScreensTimeout"
	PowerTurnOffControllersTimeoutFloat         = "turnOffControllersTimeout"
	PowerReturnToWatchdogTimeoutFloat           = "returnToWatchdogTimeout"
	PowerAutoLaunchSteamVROnButtonPress         = "autoLaunchSteamVROnButtonPress"
	PauseCompositorOnStandbyBool                = "pauseCompositorOnStandby"
	PowerPauseCompositorOnStandbyBool           = "pauseCompositorOnStandby"
	DashboardSection                            = "dashboard"
	DashboardEnableDashboardBool                = "enableDashboard"
	DashboardArcadeModeBool                     = "arcadeMode"
	ModelskinSection                            = "modelskins"
	DriverEnableBool                            = "enable"
	IVRScreenshotsVersion                       = "IVRScreenshots_001"
	IVRResourcesVersion                         = "IVRResources_001"
	IVRDriverManagerVersion                     = "IVRDriverManager_001"
)

// OpenVR Enums

// EVREye
const (
	EyeLeft  = 0
	EyeRight = 1
)

// ETextureType
const (
	TextureTypeDirectX   = 0
	TextureTypeOpenGL    = 1
	TextureTypeVulkan    = 2
	TextureTypeIOSurface = 3
	TextureTypeDirectX12 = 4
)

// EColorSpace
const (
	ColorSpaceAuto   = 0
	ColorSpaceGamma  = 1
	ColorSpaceLinear = 2
)

// ETrackingResult
const (
	TrackingResultUninitialized         = 1
	TrackingResultCalibratingInProgress = 100
	TrackingResultCalibratingOutOfRange = 101
	TrackingResultRunningOK             = 200
	TrackingResultRunningOutOfRange     = 201
)

// ETrackedDeviceClass
const (
	TrackedDeviceClassInvalid           = 0
	TrackedDeviceClassHMD               = 1
	TrackedDeviceClassController        = 2
	TrackedDeviceClassGenericTracker    = 3
	TrackedDeviceClassTrackingReference = 4
	TrackedDeviceClassDisplayRedirect   = 5
)

// ETrackedControllerRole
const (
	TrackedControllerRoleInvalid   = 0
	TrackedControllerRoleLeftHand  = 1
	TrackedControllerRoleRightHand = 2
)

// ETrackingUniverseOrigin
const (
	TrackingUniverseSeated             = 0
	TrackingUniverseStanding           = 1
	TrackingUniverseRawAndUncalibrated = 2
)

// ETrackedDeviceProperty
const (
	PropInvalid                                       = 0
	PropTrackingSystemNameString                      = 1000
	PropModelNumberString                             = 1001
	PropSerialNumberString                            = 1002
	PropRenderModelNameString                         = 1003
	PropWillDriftInYawBool                            = 1004
	PropManufacturerNameString                        = 1005
	PropTrackingFirmwareVersionString                 = 1006
	PropHardwareRevisionString                        = 1007
	PropAllWirelessDongleDescriptionsString           = 1008
	PropConnectedWirelessDongleString                 = 1009
	PropDeviceIsWirelessBool                          = 1010
	PropDeviceIsChargingBool                          = 1011
	PropDeviceBatteryPercentageFloat                  = 1012
	PropStatusDisplayTransformMatrix34                = 1013
	PropFirmwareUpdateAvailableBool                   = 1014
	PropFirmwareManualUpdateBool                      = 1015
	PropFirmwareManualUpdateURLString                 = 1016
	PropHardwareRevisionUint64                        = 1017
	PropFirmwareVersionUint64                         = 1018
	PropFPGAVersionUint64                             = 1019
	PropVRCVersionUint64                              = 1020
	PropRadioVersionUint64                            = 1021
	PropDongleVersionUint64                           = 1022
	PropBlockServerShutdownBool                       = 1023
	PropCanUnifyCoordinateSystemWithHmdBool           = 1024
	PropContainsProximitySensorBool                   = 1025
	PropDeviceProvidesBatteryStatusBool               = 1026
	PropDeviceCanPowerOffBool                         = 1027
	PropFirmwareProgrammingTargetString               = 1028
	PropDeviceClassInt32                              = 1029
	PropHasCameraBool                                 = 1030
	PropDriverVersionString                           = 1031
	PropFirmwareForceUpdateRequiredBool               = 1032
	PropViveSystemButtonFixRequiredBool               = 1033
	PropParentDriverUint64                            = 1034
	PropResourceRootString                            = 1035
	PropRegisteredDeviceTypeString                    = 1036
	PropInputProfilePathString                        = 1037
	PropNeverTrackedBool                              = 1038
	PropNumCamerasInt32                               = 1039
	PropCameraFrameLayoutInt32                        = 1040
	PropCameraStreamFormatInt32                       = 1041
	PropAdditionalDeviceSettingsPathString            = 1042
	PropIdentifiableBool                              = 1043
	PropBootloaderVersionUint64                       = 1044
	PropAdditionalSystemReportDataString              = 1045
	PropCompositeFirmwareVersionString                = 1046
	PropFirmwareRemindUpdateBool                      = 1047
	PropPeripheralApplicationVersionUint64            = 1048
	PropManufacturerSerialNumberString                = 1049
	PropComputedSerialNumberString                    = 1050
	PropEstimatedDeviceFirstUseTimeInt32              = 1051
	PropReportsTimeSinceVSyncBool                     = 2000
	PropSecondsFromVsyncToPhotonsFloat                = 2001
	PropDisplayFrequencyFloat                         = 2002
	PropUserIpdMetersFloat                            = 2003
	PropCurrentUniverseIdUint64                       = 2004
	PropPreviousUniverseIdUint64                      = 2005
	PropDisplayFirmwareVersionUint64                  = 2006
	PropIsOnDesktopBool                               = 2007
	PropDisplayMCTypeInt32                            = 2008
	PropDisplayMCOffsetFloat                          = 2009
	PropDisplayMCScaleFloat                           = 2010
	PropEdidVendorIDInt32                             = 2011
	PropDisplayMCImageLeftString                      = 2012
	PropDisplayMCImageRightString                     = 2013
	PropDisplayGCBlackClampFloat                      = 2014
	PropEdidProductIDInt32                            = 2015
	PropCameraToHeadTransformMatrix34                 = 2016
	PropDisplayGCTypeInt32                            = 2017
	PropDisplayGCOffsetFloat                          = 2018
	PropDisplayGCScaleFloat                           = 2019
	PropDisplayGCPrescaleFloat                        = 2020
	PropDisplayGCImageString                          = 2021
	PropLensCenterLeftUFloat                          = 2022
	PropLensCenterLeftVFloat                          = 2023
	PropLensCenterRightUFloat                         = 2024
	PropLensCenterRightVFloat                         = 2025
	PropUserHeadToEyeDepthMetersFloat                 = 2026
	PropCameraFirmwareVersionUint64                   = 2027
	PropCameraFirmwareDescriptionString               = 2028
	PropDisplayFPGAVersionUint64                      = 2029
	PropDisplayBootloaderVersionUint64                = 2030
	PropDisplayHardwareVersionUint64                  = 2031
	PropAudioFirmwareVersionUint64                    = 2032
	PropCameraCompatibilityModeInt32                  = 2033
	PropScreenshotHorizontalFieldOfViewDegreesFloat   = 2034
	PropScreenshotVerticalFieldOfViewDegreesFloat     = 2035
	PropDisplaySuppressedBool                         = 2036
	PropDisplayAllowNightModeBool                     = 2037
	PropDisplayMCImageWidthInt32                      = 2038
	PropDisplayMCImageHeightInt32                     = 2039
	PropDisplayMCImageNumChannelsInt32                = 2040
	PropDisplayMCImageDataBinary                      = 2041
	PropSecondsFromPhotonsToVblankFloat               = 2042
	PropDriverDirectModeSendsVsyncEventsBool          = 2043
	PropDisplayDebugModeBool                          = 2044
	PropGraphicsAdapterLuidUint64                     = 2045
	PropDriverProvidedChaperonePathString             = 2048
	PropExpectedTrackingReferenceCountInt32           = 2049
	PropExpectedControllerCountInt32                  = 2050
	PropNamedIconPathControllerLeftDeviceOffString    = 2051
	PropNamedIconPathControllerRightDeviceOffString   = 2052
	PropNamedIconPathTrackingReferenceDeviceOffString = 2053
	PropDoNotApplyPredictionBool                      = 2054
	PropCameraToHeadTransformsMatrix34Array           = 2055
	PropDistortionMeshResolutionInt32                 = 2056
	PropDriverIsDrawingControllersBool                = 2057
	PropDriverRequestsApplicationPauseBool            = 2058
	PropDriverRequestsReducedRenderingBool            = 2059
	PropMinimumIpdStepMetersFloat                     = 2060
	PropAudioBridgeFirmwareVersionUint64              = 2061
	PropImageBridgeFirmwareVersionUint64              = 2062
	PropImuToHeadTransformMatrix34                    = 2063
	PropImuFactoryGyroBiasVector3                     = 2064
	PropImuFactoryGyroScaleVector3                    = 2065
	PropImuFactoryAccelerometerBiasVector3            = 2066
	PropImuFactoryAccelerometerScaleVector3           = 2067
	PropConfigurationIncludesLighthouse20FeaturesBool = 2069
	PropAdditionalRadioFeaturesUint64                 = 2070
	PropCameraWhiteBalanceVector4Array                = 2071
	PropCameraDistortionFunctionInt32Array            = 2072
	PropCameraDistortionCoefficientsFloatArray        = 2073
	PropExpectedControllerTypeString                  = 2074
	PropHmdTrackingStyleInt32                         = 2075
	PropDriverProvidedChaperoneVisibilityBool         = 2076
	PropHmdColumnCorrectionSettingPrefixString        = 2077
	PropCameraSupportsCompatibilityModesBool          = 2078
	PropSupportsRoomViewDepthProjectionBool           = 2079
	PropDisplayAvailableFrameRatesFloatArray          = 2080
	PropDisplaySupportsMultipleFrameratesBool         = 2081
	PropDisplayColorMultLeftVector3                   = 2082
	PropDisplayColorMultRightVector3                  = 2083
	PropDisplaySupportsRuntimeFramerateChangeBool     = 2084
	PropDisplaySupportsAnalogGainBool                 = 2085
	PropDisplayMinAnalogGainFloat                     = 2086
	PropDisplayMaxAnalogGainFloat                     = 2087
	PropCameraExposureTimeFloat                       = 2088
	PropCameraGlobalGainFloat                         = 2089
	PropDashboardScaleFloat                           = 2091
	PropIpdUIRangeMinMetersFloat                      = 2100
	PropIpdUIRangeMaxMetersFloat                      = 2101
	PropHmdSupportsHDCP14LegacyCompatBool             = 2102
	PropHmdSupportsMicMonitoringBool                  = 2103
	PropDriverRequestedMuraCorrectionModeInt32        = 2200
	PropDriverRequestedMuraFeatherInnerLeftInt32      = 2201
	PropDriverRequestedMuraFeatherInnerRightInt32     = 2202
	PropDriverRequestedMuraFeatherInnerTopInt32       = 2203
	PropDriverRequestedMuraFeatherInnerBottomInt32    = 2204
	PropDriverRequestedMuraFeatherOuterLeftInt32      = 2205
	PropDriverRequestedMuraFeatherOuterRightInt32     = 2206
	PropDriverRequestedMuraFeatherOuterTopInt32       = 2207
	PropDriverRequestedMuraFeatherOuterBottomInt32    = 2208
	PropAudioDefaultPlaybackDeviceIdString            = 2300
	PropAudioDefaultRecordingDeviceIdString           = 2301
	PropAudioDefaultPlaybackDeviceVolumeFloat         = 2302
	PropAudioSupportsDualSpeakerAndJackOutputBool     = 2303
	PropAttachedDeviceIdString                        = 3000
	PropSupportedButtonsUint64                        = 3001
	PropAxis0TypeInt32                                = 3002
	PropAxis1TypeInt32                                = 3003
	PropAxis2TypeInt32                                = 3004
	PropAxis3TypeInt32                                = 3005
	PropAxis4TypeInt32                                = 3006
	PropControllerRoleHintInt32                       = 3007
	PropFieldOfViewLeftDegreesFloat                   = 4000
	PropFieldOfViewRightDegreesFloat                  = 4001
	PropFieldOfViewTopDegreesFloat                    = 4002
	PropFieldOfViewBottomDegreesFloat                 = 4003
	PropTrackingRangeMinimumMetersFloat               = 4004
	PropTrackingRangeMaximumMetersFloat               = 4005
	PropModeLabelString                               = 4006
	PropCanWirelessIdentifyBool                       = 4007
	PropNonceInt32                                    = 4008
	PropIconPathNameString                            = 5000
	PropNamedIconPathDeviceOffString                  = 5001
	PropNamedIconPathDeviceSearchingString            = 5002
	PropNamedIconPathDeviceSearchingAlertString       = 5003
	PropNamedIconPathDeviceReadyString                = 5004
	PropNamedIconPathDeviceReadyAlertString           = 5005
	PropNamedIconPathDeviceNotReadyString             = 5006
	PropNamedIconPathDeviceStandbyString              = 5007
	PropNamedIconPathDeviceAlertLowString             = 5008
	PropNamedIconPathDeviceStandbyAlertString         = 5009
	PropDisplayHiddenAreaBinaryStart                  = 5100
	PropDisplayHiddenAreaBinaryEnd                    = 5150
	PropParentContainer                               = 5151
	PropOverrideContainerUint64                       = 5152
	PropUserConfigPathString                          = 6000
	PropInstallPathString                             = 6001
	PropHasDisplayComponentBool                       = 6002
	PropHasControllerComponentBool                    = 6003
	PropHasCameraComponentBool                        = 6004
	PropHasDriverDirectModeComponentBool              = 6005
	PropHasVirtualDisplayComponentBool                = 6006
	PropHasSpatialAnchorsSupportBool                  = 6007
	PropControllerTypeString                          = 7000
	PropControllerHandSelectionPriorityInt32          = 7002
	PropVendorSpecificReservedStart                   = 10000
	PropVendorSpecificReservedEnd                     = 10999
	PropTrackedDevicePropertyMax                      = 1000000
)

// ETrackedPropertyError
const (
	TrackedPropSuccess                    = 0
	TrackedPropWrongDataType              = 1
	TrackedPropWrongDeviceClass           = 2
	TrackedPropBufferTooSmall             = 3
	TrackedPropUnknownProperty            = 4
	TrackedPropInvalidDevice              = 5
	TrackedPropCouldNotContactServer      = 6
	TrackedPropValueNotProvidedByDevice   = 7
	TrackedPropStringExceedsMaximumLength = 8
	TrackedPropNotYetAvailable            = 9
	TrackedPropPermissionDenied           = 10
	TrackedPropInvalidOperation           = 11
)

// EVRSubmitFlags
const (
	SubmitDefault                      = 0
	SubmitLensDistortionAlreadyApplied = 1
	SubmitGlRenderBuffer               = 2
	SubmitReserved                     = 4
	SubmitTextureWithPose              = 8
)

// EVRState
const (
	VRStateUndefined      = -1
	VRStateOff            = 0
	VRStateSearching      = 1
	VRStateSearchingAlert = 2
	VRStateReady          = 3
	VRStateReadyAlert     = 4
	VRStateNotReady       = 5
	VRStateStandby        = 6
	VRStateReadyAlertLow  = 7
)

// EVREventType
const (
	VREventNone                                      = 0
	VREventTrackedDeviceActivated                    = 100
	VREventTrackedDeviceDeactivated                  = 101
	VREventTrackedDeviceUpdated                      = 102
	VREventTrackedDeviceUserInteractionStarted       = 103
	VREventTrackedDeviceUserInteractionEnded         = 104
	VREventIpdChanged                                = 105
	VREventEnterStandbyMode                          = 106
	VREventLeaveStandbyMode                          = 107
	VREventTrackedDeviceRoleChanged                  = 108
	VREventWatchdogWakeUpRequested                   = 109
	VREventLensDistortionChanged                     = 110
	VREventPropertyChanged                           = 111
	VREventWirelessDisconnect                        = 112
	VREventWirelessReconnect                         = 113
	VREventButtonPress                               = 200
	VREventButtonUnpress                             = 201
	VREventButtonTouch                               = 202
	VREventButtonUntouch                             = 203
	VREventModalCancel                               = 257
	VREventMouseMove                                 = 300
	VREventMouseButtonDown                           = 301
	VREventMouseButtonUp                             = 302
	VREventFocusEnter                                = 303
	VREventFocusLeave                                = 304
	VREventScrollDiscrete                            = 305
	VREventTouchPadMove                              = 306
	VREventOverlayFocusChanged                       = 307
	VREventReloadOverlays                            = 308
	VREventScrollSmooth                              = 309
	VREventLockMousePosition                         = 310
	VREventUnlockMousePosition                       = 311
	VREventInputFocusCaptured                        = 400
	VREventInputFocusReleased                        = 401
	VREventSceneApplicationChanged                   = 404
	VREventSceneFocusChanged                         = 405
	VREventInputFocusChanged                         = 406
	VREventSceneApplicationUsingWrongGraphicsAdapter = 408
	VREventActionBindingReloaded                     = 409
	VREventHideRenderModels                          = 410
	VREventShowRenderModels                          = 411
	VREventSceneApplicationStateChanged              = 412
	VREventConsoleOpened                             = 420
	VREventConsoleClosed                             = 421
	VREventOverlayShown                              = 500
	VREventOverlayHidden                             = 501
	VREventDashboardActivated                        = 502
	VREventDashboardDeactivated                      = 503
	VREventDashboardRequested                        = 505
	VREventResetDashboard                            = 506
	VREventImageLoaded                               = 508
	VREventShowKeyboard                              = 509
	VREventHideKeyboard                              = 510
	VREventOverlayGamepadFocusGained                 = 511
	VREventOverlayGamepadFocusLost                   = 512
	VREventOverlaySharedTextureChanged               = 513
	VREventScreenshotTriggered                       = 516
	VREventImageFailed                               = 517
	VREventDashboardOverlayCreated                   = 518
	VREventSwitchGamepadFocus                        = 519
	VREventRequestScreenshot                         = 520
	VREventScreenshotTaken                           = 521
	VREventScreenshotFailed                          = 522
	VREventSubmitScreenshotToDashboard               = 523
	VREventScreenshotProgressToDashboard             = 524
	VREventPrimaryDashboardDeviceChanged             = 525
	VREventRoomViewShown                             = 526
	VREventRoomViewHidden                            = 527
	VREventShowUI                                    = 528
	VREventShowDevTools                              = 529
	VREventDesktopViewUpdating                       = 530
	VREventDesktopViewReady                          = 531
	VREventNotificationShown                         = 600
	VREventNotificationHidden                        = 601
	VREventNotificationBeginInteraction              = 602
	VREventNotificationDestroyed                     = 603
	VREventQuit                                      = 700
	VREventProcessQuit                               = 701
	VREventQuitAcknowledged                          = 703
	VREventDriverRequestedQuit                       = 704
	VREventRestartRequested                          = 705
	VREventChaperoneDataHasChanged                   = 800
	VREventChaperoneUniverseHasChanged               = 801
	VREventChaperoneTempDataHasChanged               = 802
	VREventChaperoneSettingsHaveChanged              = 803
	VREventSeatedZeroPoseReset                       = 804
	VREventChaperoneFlushCache                       = 805
	VREventChaperoneRoomSetupStarting                = 806
	VREventChaperoneRoomSetupFinished                = 807
	VREventStandingZeroPoseReset                     = 808
	VREventAudioSettingsHaveChanged                  = 820
	VREventBackgroundSettingHasChanged               = 850
	VREventCameraSettingsHaveChanged                 = 851
	VREventReprojectionSettingHasChanged             = 852
	VREventModelSkinSettingsHaveChanged              = 853
	VREventEnvironmentSettingsHaveChanged            = 854
	VREventPowerSettingsHaveChanged                  = 855
	VREventEnableHomeAppSettingsHaveChanged          = 856
	VREventSteamVRSectionSettingChanged              = 857
	VREventLighthouseSectionSettingChanged           = 858
	VREventNullSectionSettingChanged                 = 859
	VREventUserInterfaceSectionSettingChanged        = 860
	VREventNotificationsSectionSettingChanged        = 861
	VREventKeyboardSectionSettingChanged             = 862
	VREventPerfSectionSettingChanged                 = 863
	VREventDashboardSectionSettingChanged            = 864
	VREventWebInterfaceSectionSettingChanged         = 865
	VREventTrackersSectionSettingChanged             = 866
	VREventLastKnownSectionSettingChanged            = 867
	VREventDismissedWarningsSectionSettingChanged    = 868
	VREventGpuSpeedSectionSettingChanged             = 869
	VREventWindowsMRSectionSettingChanged            = 870
	VREventOtherSectionSettingChanged                = 871
	VREventStatusUpdate                              = 900
	VREventWebInterfaceInstallDriverCompleted        = 950
	VREventMCImageUpdated                            = 1000
	VREventFirmwareUpdateStarted                     = 1100
	VREventFirmwareUpdateFinished                    = 1101
	VREventKeyboardClosed                            = 1200
	VREventKeyboardCharInput                         = 1201
	VREventKeyboardDone                              = 1202
	VREventApplicationListUpdated                    = 1303
	VREventApplicationMimeTypeLoad                   = 1304
	VREventProcessConnected                          = 1306
	VREventProcessDisconnected                       = 1307
	VREventCompositorChaperoneBoundsShown            = 1410
	VREventCompositorChaperoneBoundsHidden           = 1411
	VREventCompositorDisplayDisconnected             = 1412
	VREventCompositorDisplayReconnected              = 1413
	VREventCompositorHDCPError                       = 1414
	VREventCompositorApplicationNotResponding        = 1415
	VREventCompositorApplicationResumed              = 1416
	VREventCompositorOutOfVideoMemory                = 1417
	VREventCompositorDisplayModeNotSupported         = 1418
	VREventCompositorStageOverrideReady              = 1419
	VREventTrackedCameraStartVideoStream             = 1500
	VREventTrackedCameraStopVideoStream              = 1501
	VREventTrackedCameraPauseVideoStream             = 1502
	VREventTrackedCameraResumeVideoStream            = 1503
	VREventTrackedCameraEditingSurface               = 1550
	VREventPerformanceTestEnableCapture              = 1600
	VREventPerformanceTestDisableCapture             = 1601
	VREventPerformanceTestFidelityLevel              = 1602
	VREventMessageOverlayClosed                      = 1650
	VREventMessageOverlayCloseRequested              = 1651
	VREventInputHapticVibration                      = 1700
	VREventInputBindingLoadFailed                    = 1701
	VREventInputBindingLoadSuccessful                = 1702
	VREventInputActionManifestReloaded               = 1703
	VREventInputActionManifestLoadFailed             = 1704
	VREventInputProgressUpdate                       = 1705
	VREventInputTrackerActivated                     = 1706
	VREventInputBindingsUpdated                      = 1707
	VREventInputBindingSubscriptionChanged           = 1708
	VREventSpatialAnchorsPoseUpdated                 = 1800
	VREventSpatialAnchorsDescriptorUpdated           = 1801
	VREventSpatialAnchorsRequestPoseUpdate           = 1802
	VREventSpatialAnchorsRequestDescriptorUpdate     = 1803
	VREventSystemReportStarted                       = 1900
	VREventMonitorShowHeadsetView                    = 2000
	VREventMonitorHideHeadsetView                    = 2001
	VREventVendorSpecificReservedStart               = 10000
	VREventVendorSpecificReservedEnd                 = 19999
)

// EDeviceActivityLevel
const (
	DeviceActivityLevelUnknown                = -1
	DeviceActivityLevelIdle                   = 0
	DeviceActivityLevelUserInteraction        = 1
	DeviceActivityLevelUserInteractionTimeout = 2
	DeviceActivityLevelStandby                = 3
)

// EVRButtonId
const (
	ButtonSystem          = 0
	ButtonApplicationMenu = 1
	ButtonGrip            = 2
	ButtonDPadLeft        = 3
	ButtonDPadUp          = 4
	ButtonDPadRight       = 5
	ButtonDPadDown        = 6
	ButtonA               = 7
	ButtonProximitySensor = 31
	ButtonAxis0           = 32
	ButtonAxis1           = 33
	ButtonAxis2           = 34
	ButtonAxis3           = 35
	ButtonAxis4           = 36
	ButtonSteamVRTouchpad = 32
	ButtonSteamVRTrigger  = 33
	ButtonDashboardBack   = 2
	ButtonMax             = 64
)

// EVRMouseButton
const (
	VRMouseButtonLeft   = 1
	VRMouseButtonRight  = 2
	VRMouseButtonMiddle = 4
)

// EHiddenAreaMeshType
const (
	HiddenAreaMeshStandard = 0
	HiddenAreaMeshInverse  = 1
	HiddenAreaMeshLineLoop = 2
	HiddenAreaMeshMax      = 3
)

// EVRControllerAxisType
const (
	VRControllerAxisNone     = 0
	VRControllerAxisTrackPad = 1
	VRControllerAxisJoystick = 2
	VRControllerAxisTrigger  = 3
)

// EVRControllerEventOutputType
const (
	VRControllerEventOutputOSEvents = 0
	VRControllerEventOutputVREvents = 1
)

// ECollisionBoundsStyle
const (
	CollisionBoundsStyleBeginner     = 0
	CollisionBoundsStyleIntermediate = 1
	CollisionBoundsStyleSquares      = 2
	CollisionBoundsStyleAdvanced     = 3
	CollisionBoundsStyleNone         = 4
	CollisionBoundsStyleCount        = 5
)

// EVROverlayError
const (
	VROverlayErrorNone                     = 0
	VROverlayErrorUnknownOverlay           = 10
	VROverlayErrorInvalidHandle            = 11
	VROverlayErrorPermissionDenied         = 12
	VROverlayErrorOverlayLimitExceeded     = 13
	VROverlayErrorWrongVisibilityType      = 14
	VROverlayErrorKeyTooLong               = 15
	VROverlayErrorNameTooLong              = 16
	VROverlayErrorKeyInUse                 = 17
	VROverlayErrorWrongTransformType       = 18
	VROverlayErrorInvalidTrackedDevice     = 19
	VROverlayErrorInvalidParameter         = 20
	VROverlayErrorThumbnailCantBeDestroyed = 21
	VROverlayErrorArrayTooSmall            = 22
	VROverlayErrorRequestFailed            = 23
	VROverlayErrorInvalidTexture           = 24
	VROverlayErrorUnableToLoadFile         = 25
	VROVerlayErrorKeyboardAlreadyInUse     = 26
	VROverlayErrorNoNeighbor               = 27
	VROverlayErrorTooManyMaskPrimitives    = 29
	VROverlayErrorBadMaskPrimitive         = 30
)

// EVRApplicationType
const (
	VRApplicationOther         = 0
	VRApplicationScene         = 1
	VRApplicationOverlay       = 2
	VRApplicationBackground    = 3
	VRApplicationUtility       = 4
	VRApplicationVRMonitor     = 5
	VRApplicationSteamWatchdog = 6
	VRApplicationBootstrapper  = 7
	VRApplicationMax           = 8
)

// EVRFirmwareError
const (
	VRFirmwareErrorNone    = 0
	VRFirmwareErrorSuccess = 1
	VRFirmwareErrorFail    = 2
)

// EVRNotificationError
const (
	VRNotificationErrorOK                               = 0
	VRNotificationErrorInvalidNotificationId            = 100
	VRNotificationErrorNotificationQueueFull            = 101
	VRNotificationErrorInvalidOverlayHandle             = 102
	VRNotificationErrorSystemWithUserValueAlreadyExists = 103
)

// EVRInitError
const (
	VRInitErrorNone                                             = 0
	VRInitErrorUnknown                                          = 1
	VRInitErrorInitInstallationNotFound                         = 100
	VRInitErrorInitInstallationCorrupt                          = 101
	VRInitErrorInitVRClientDLLNotFound                          = 102
	VRInitErrorInitFileNotFound                                 = 103
	VRInitErrorInitFactoryNotFound                              = 104
	VRInitErrorInitInterfaceNotFound                            = 105
	VRInitErrorInitInvalidInterface                             = 106
	VRInitErrorInitUserConfigDirectoryInvalid                   = 107
	VRInitErrorInitHmdNotFound                                  = 108
	VRInitErrorInitNotInitialized                               = 109
	VRInitErrorInitPathRegistryNotFound                         = 110
	VRInitErrorInitNoConfigPath                                 = 111
	VRInitErrorInitNoLogPath                                    = 112
	VRInitErrorInitPathRegistryNotWritable                      = 113
	VRInitErrorInitAppInfoInitFailed                            = 114
	VRInitErrorInitRetry                                        = 115
	VRInitErrorInitInitCanceledByUser                           = 116
	VRInitErrorInitAnotherAppLaunching                          = 117
	VRInitErrorInitSettingsInitFailed                           = 118
	VRInitErrorInitShuttingDown                                 = 119
	VRInitErrorInitTooManyObjects                               = 120
	VRInitErrorInitNoServerForBackgroundApp                     = 121
	VRInitErrorInitNotSupportedWithCompositor                   = 122
	VRInitErrorInitNotAvailableToUtilityApps                    = 123
	VRInitErrorInitInternal                                     = 124
	VRInitErrorInitHmdDriverIdIsNone                            = 125
	VRInitErrorInitHmdNotFoundPresenceFailed                    = 126
	VRInitErrorInitVRMonitorNotFound                            = 127
	VRInitErrorInitVRMonitorStartupFailed                       = 128
	VRInitErrorInitLowPowerWatchdogNotSupported                 = 129
	VRInitErrorInitInvalidApplicationType                       = 130
	VRInitErrorInitNotAvailableToWatchdogApps                   = 131
	VRInitErrorInitWatchdogDisabledInSettings                   = 132
	VRInitErrorInitVRDashboardNotFound                          = 133
	VRInitErrorInitVRDashboardStartupFailed                     = 134
	VRInitErrorInitVRHomeNotFound                               = 135
	VRInitErrorInitVRHomeStartupFailed                          = 136
	VRInitErrorInitRebootingBusy                                = 137
	VRInitErrorInitFirmwareUpdateBusy                           = 138
	VRInitErrorInitFirmwareRecoveryBusy                         = 139
	VRInitErrorDriverFailed                                     = 200
	VRInitErrorDriverUnknown                                    = 201
	VRInitErrorDriverHmdUnknown                                 = 202
	VRInitErrorDriverNotLoaded                                  = 203
	VRInitErrorDriverRuntimeOutOfDate                           = 204
	VRInitErrorDriverHmdInUse                                   = 205
	VRInitErrorDriverNotCalibrated                              = 206
	VRInitErrorDriverCalibrationInvalid                         = 207
	VRInitErrorDriverHmdDisplayNotFound                         = 208
	VRInitErrorDriverTrackedDeviceInterfaceUnknown              = 209
	VRInitErrorDriverHmdDriverIdOutOfBounds                     = 211
	VRInitErrorDriverHmdDisplayMirrored                         = 212
	VRInitErrorIPCServerInitFailed                              = 300
	VRInitErrorIPCConnectFailed                                 = 301
	VRInitErrorIPCSharedStateInitFailed                         = 302
	VRInitErrorIPCCompositorInitFailed                          = 303
	VRInitErrorIPCMutexInitFailed                               = 304
	VRInitErrorIPCFailed                                        = 305
	VRInitErrorIPCCompositorConnectFailed                       = 306
	VRInitErrorIPCCompositorInvalidConnectResponse              = 307
	VRInitErrorIPCConnectFailedAfterMultipleAttempts            = 308
	VRInitErrorCompositorFailed                                 = 400
	VRInitErrorCompositorD3D11HardwareRequired                  = 401
	VRInitErrorCompositorFirmwareRequiresUpdate                 = 402
	VRInitErrorCompositorOverlayInitFailed                      = 403
	VRInitErrorCompositorScreenshotsInitFailed                  = 404
	VRInitErrorCompositorUnableToCreateDevice                   = 405
	VRInitErrorVendorSpecificUnableToConnectToOculusRuntime     = 1000
	VRInitErrorVendorSpecificHmdFoundCantOpenDevice             = 1101
	VRInitErrorVendorSpecificHmdFoundUnableToRequestConfigStart = 1102
	VRInitErrorVendorSpecificHmdFoundNoStoredConfig             = 1103
	VRInitErrorVendorSpecificHmdFoundConfigTooBig               = 1104
	VRInitErrorVendorSpecificHmdFoundConfigTooSmall             = 1105
	VRInitErrorVendorSpecificHmdFoundUnableToInitZLib           = 1106
	VRInitErrorVendorSpecificHmdFoundCantReadFirmwareVersion    = 1107
	VRInitErrorVendorSpecificHmdFoundUnableToSendUserDataStart  = 1108
	VRInitErrorVendorSpecificHmdFoundUnableToGetUserDataStart   = 1109
	VRInitErrorVendorSpecificHmdFoundUnableToGetUserDataNext    = 1110
	VRInitErrorVendorSpecificHmdFoundUserDataAddressRange       = 1111
	VRInitErrorVendorSpecificHmdFoundUserDataError              = 1112
	VRInitErrorVendorSpecificHmdFoundConfigFailedSanityCheck    = 1113
	VRInitErrorSteamSteamInstallationNotFound                   = 2000
)

// EVRScreenshotType
const (
	VRScreenshotTypeNone           = 0
	VRScreenshotTypeMono           = 1
	VRScreenshotTypeStereo         = 2
	VRScreenshotTypeCubemap        = 3
	VRScreenshotTypeMonoPanorama   = 4
	VRScreenshotTypeStereoPanorama = 5
)

// EVRScreenshotPropertyFilenames
const (
	VRScreenshotPropertyFilenamesPreview = 0
	VRScreenshotPropertyFilenamesVR      = 1
)

// EVRTrackedCameraError
const (
	VRTrackedCameraErrorNone                       = 0
	VRTrackedCameraErrorOperationFailed            = 100
	VRTrackedCameraErrorInvalidHandle              = 101
	VRTrackedCameraErrorInvalidFrameHeaderVersion  = 102
	VRTrackedCameraErrorOutOfHandles               = 103
	VRTrackedCameraErrorIPCFailure                 = 104
	VRTrackedCameraErrorNotSupportedForThisDevice  = 105
	VRTrackedCameraErrorSharedMemoryFailure        = 106
	VRTrackedCameraErrorFrameBufferingFailure      = 107
	VRTrackedCameraErrorStreamSetupFailure         = 108
	VRTrackedCameraErrorInvalidGLTextureId         = 109
	VRTrackedCameraErrorInvalidSharedTextureHandle = 110
	VRTrackedCameraErrorFailedToGetGLTextureId     = 111
	VRTrackedCameraErrorSharedTextureFailure       = 112
	VRTrackedCameraErrorNoFrameAvailable           = 113
	VRTrackedCameraErrorInvalidArgument            = 114
	VRTrackedCameraErrorInvalidFrameBufferSize     = 115
)

// EVRTrackedCameraFrameType
const (
	VRTrackedCameraFrameTypeDistorted          = 0
	VRTrackedCameraFrameTypeUndistorted        = 1
	VRTrackedCameraFrameTypeMaximumUndistorted = 2
	VRTrackedCameraMaxCameraFrameTypes         = 3
)

// EVRApplicationError
const (
	VRApplicationErrorNone                       = 0
	VRApplicationErrorAppKeyAlreadyExists        = 100
	VRApplicationErrorNoManifest                 = 101
	VRApplicationErrorNoApplication              = 102
	VRApplicationErrorInvalidIndex               = 103
	VRApplicationErrorUnknownApplication         = 104
	VRApplicationErrorIPCFailed                  = 105
	VRApplicationErrorApplicationAlreadyRunning  = 106
	VRApplicationErrorInvalidManifest            = 107
	VRApplicationErrorInvalidApplication         = 108
	VRApplicationErrorLaunchFailed               = 109
	VRApplicationErrorApplicationAlreadyStarting = 110
	VRApplicationErrorLaunchInProgress           = 111
	VRApplicationErrorOldApplicationQuitting     = 112
	VRApplicationErrorTransitionAborted          = 113
	VRApplicationErrorIsTemplate                 = 114
	VRApplicationErrorSteamVRIsExiting           = 115
	VRApplicationErrorBufferTooSmall             = 200
	VRApplicationErrorPropertyNotSet             = 201
	VRApplicationErrorUnknownProperty            = 202
	VRApplicationErrorInvalidParameter           = 203
)

// EVRApplicationProperty
const (
	VRApplicationPropertyNameString                        = 0
	VRApplicationPropertyLaunchTypeString                  = 11
	VRApplicationPropertyWorkingDirectoryString            = 12
	VRApplicationPropertyBinaryPathString                  = 13
	VRApplicationPropertyArgumentsString                   = 14
	VRApplicationPropertyURLString                         = 15
	VRApplicationPropertyDescriptionString                 = 50
	VRApplicationPropertyNewsURLString                     = 51
	VRApplicationPropertyImagePathString                   = 52
	VRApplicationPropertySourceString                      = 53
	VRApplicationPropertyIsDashboardOverlayBool            = 60
	VRApplicationPropertyIsTemplateBool                    = 61
	VRApplicationPropertyIsInstancedBool                   = 62
	VRApplicationPropertyIsInternalBool                    = 63
	VRApplicationPropertyWantsCompositorPauseInStandbyBool = 64
	VRApplicationPropertyLastLaunchTimeUint64              = 70
)

// EVRApplicationTransitionState
const (
	VRApplicationTransitionNone                     = 0
	VRApplicationTransitionOldAppQuitSent           = 10
	VRApplicationTransitionWaitingForExternalLaunch = 11
	VRApplicationTransitionNewAppLaunched           = 20
)

// ChaperoneCalibrationState
const (
	ChaperoneCalibrationStateOK                             = 1
	ChaperoneCalibrationStateWarning                        = 100
	ChaperoneCalibrationStateWarningBaseStationMayHaveMoved = 101
	ChaperoneCalibrationStateWarningBaseStationRemoved      = 102
	ChaperoneCalibrationStateWarningSeatedBoundsInvalid     = 103
	ChaperoneCalibrationStateError                          = 200
	ChaperoneCalibrationStateErrorBaseStationUninitialized  = 201
	ChaperoneCalibrationStateErrorBaseStationConflict       = 202
	ChaperoneCalibrationStateErrorPlayAreaInvalid           = 203
	ChaperoneCalibrationStateErrorCollisionBoundsInvalid    = 204
)

// EChaperoneConfigFile
const (
	ChaperoneConfigFileLive = 1
	ChaperoneConfigFileTemp = 2
)

// EChaperoneImportFlags
const (
	ChaperoneImportBoundsOnly = 1
)

// EVRCompositorError
const (
	VRCompositorErrorNone                         = 0
	VRCompositorErrorRequestFailed                = 1
	VRCompositorErrorIncompatibleVersion          = 100
	VRCompositorErrorDoNotHaveFocus               = 101
	VRCompositorErrorInvalidTexture               = 102
	VRCompositorErrorIsNotSceneApplication        = 103
	VRCompositorErrorTextureIsOnWrongDevice       = 104
	VRCompositorErrorTextureUsesUnsupportedFormat = 105
	VRCompositorErrorSharedTexturesNotSupported   = 106
	VRCompositorErrorIndexOutOfRange              = 107
	VRCompositorErrorAlreadySubmitted             = 108
	VRCompositorErrorInvalidBounds                = 109
)

// VROverlayInputMethod
const (
	VROverlayInputMethodNone  = 0
	VROverlayInputMethodMouse = 1
)

// VROverlayTransformType
const (
	VROverlayTransformAbsolute              = 0
	VROverlayTransformTrackedDeviceRelative = 1
	VROverlayTransformSystemOverlay         = 2
	VROverlayTransformTrackedComponent      = 3
)

// VROverlayFlags
const (
	VROverlayFlagsNone                               = 0
	VROverlayFlagsCurved                             = 1
	VROverlayFlagsRGSS4X                             = 2
	VROverlayFlagsNoDashboardTab                     = 3
	VROverlayFlagsAcceptsGamepadEvents               = 4
	VROverlayFlagsShowGamepadFocus                   = 5
	VROverlayFlagsSendVRScrollEvents                 = 6
	VROverlayFlagsSendVRTouchpadEvents               = 7
	VROverlayFlagsShowTouchPadScrollWheel            = 8
	VROverlayFlagsTransferOwnershipToInternalProcess = 9
	VROverlayFlagsSideBySideParallel                 = 10
	VROverlayFlagsSideBySideCrossed                  = 11
	VROverlayFlagsPanorama                           = 12
	VROverlayFlagsStereoPanorama                     = 13
	VROverlayFlagsSortWithNonSceneOverlays           = 14
	VROverlayFlagsVisibleInDashboard                 = 15
)

// VRMessageOverlayResponse
const (
	VRMessageOverlayResponseButtonPress0                     = 0
	VRMessageOverlayResponseButtonPress1                     = 1
	VRMessageOverlayResponseButtonPress2                     = 2
	VRMessageOverlayResponseButtonPress3                     = 3
	VRMessageOverlayResponseCouldntFindSystemOverlay         = 4
	VRMessageOverlayResponseCouldntFindOrCreateClientOverlay = 5
	VRMessageOverlayResponseApplicationQuit                  = 6
)

// EGamepadTextInputMode
const (
	GamepadTextInputModeNormal   = 0
	GamepadTextInputModePassword = 1
	GamepadTextInputModeSubmit   = 2
)

// EGamepadTextInputLineMode
const (
	GamepadTextInputLineModeSingleLine    = 0
	GamepadTextInputLineModeMultipleLines = 1
)

// EOverlayDirection
const (
	OverlayDirectionUp    = 0
	OverlayDirectionDown  = 1
	OverlayDirectionLeft  = 2
	OverlayDirectionRight = 3
	OverlayDirectionCount = 4
)

// EVROverlayIntersectionMaskPrimitiveType
const (
	VROverlayIntersectionMaskPrimitiveTypeRectangle = 0
	VROverlayIntersectionMaskPrimitiveTypeCircle    = 1
)

// EVRRenderModelError
const (
	VRRenderModelErrorNone               = 0
	VRRenderModelErrorLoading            = 100
	VRRenderModelErrorNotSupported       = 200
	VRRenderModelErrorInvalidArg         = 300
	VRRenderModelErrorInvalidModel       = 301
	VRRenderModelErrorNoShapes           = 302
	VRRenderModelErrorMultipleShapes     = 303
	VRRenderModelErrorTooManyVertices    = 304
	VRRenderModelErrorMultipleTextures   = 305
	VRRenderModelErrorBufferTooSmall     = 306
	VRRenderModelErrorNotEnoughNormals   = 307
	VRRenderModelErrorNotEnoughTexCoords = 308
	VRRenderModelErrorInvalidTexture     = 400
)

// EVRComponentProperty
const (
	VRComponentPropertyIsStatic   = 1
	VRComponentPropertyIsVisible  = 2
	VRComponentPropertyIsTouched  = 4
	VRComponentPropertyIsPressed  = 8
	VRComponentPropertyIsScrolled = 16
)

// EVRNotificationType
const (
	VRNotificationTypeTransient                    = 0
	VRNotificationTypePersistent                   = 1
	VRNotificationTypeTransientSystemWithUserValue = 2
)

// EVRNotificationStyle
const (
	VRNotificationStyleNone            = 0
	VRNotificationStyleApplication     = 100
	VRNotificationStyleContactDisabled = 200
	VRNotificationStyleContactEnabled  = 201
	VRNotificationStyleContactActive   = 202
)

// EVRSettingsError
const (
	VRSettingsErrorNone                     = 0
	VRSettingsErrorIPCFailed                = 1
	VRSettingsErrorWriteFailed              = 2
	VRSettingsErrorReadFailed               = 3
	VRSettingsErrorJsonParseFailed          = 4
	VRSettingsErrorUnsetSettingHasNoDefault = 5
)

// EVRScreenshotError
const (
	VRScreenshotErrorNone                        = 0
	VRScreenshotErrorRequestFailed               = 1
	VRScreenshotErrorIncompatibleVersion         = 100
	VRScreenshotErrorNotFound                    = 101
	VRScreenshotErrorBufferTooSmall              = 102
	VRScreenshotErrorScreenshotAlreadyInProgress = 108
)
