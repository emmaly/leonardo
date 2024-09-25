package leonardo

import (
	"fmt"
	"strings"
	"time"
)

// APIErrorResponse represents a detailed error response from the API.
type APIErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"error"`
	Path    string `json:"path"`
}

// Helper function to create pointers
func Ptr[T any](v T) *T {
	return &v
}

type Time struct {
	Time time.Time
}

func (t *Time) String() string {
	return t.Time.Format("2006-01-02T15:04:05.000")
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// Remove quotes from string
	s := strings.Trim(string(data), "\"")

	// Parse the time string
	pt, err := time.Parse("2006-01-02T15:04:05.000", s)
	if err != nil {
		return err
	}

	t.Time = pt
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format("2006-01-02T15:04:05.000"))), nil
}

// Dataset-related types

// CreateDatasetRequest represents the payload for creating a new dataset.
type CreateDatasetRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// CreateDatasetResponse represents the response after creating a dataset.
type CreateDatasetResponse struct {
	InsertDatasetsOne struct {
		ID *string `json:"id"`
	} `json:"insert_datasets_one"`
}

// GetDatasetResponse represents the response when retrieving a dataset by ID.
type GetDatasetResponse struct {
	DatasetsByPk struct {
		CreatedAt     *Time `json:"createdAt"`
		DatasetImages []struct {
			CreatedAt *Time   `json:"createdAt"`
			ID        *string `json:"id"`
			URL       *string `json:"url"`
		} `json:"dataset_images"`
		Description *string `json:"description"`
		ID          *string `json:"id"`
		Name        *string `json:"name"`
		UpdatedAt   *Time   `json:"updatedAt"`
	} `json:"datasets_by_pk"`
}

// DeleteDatasetResponse represents the response after deleting a dataset.
type DeleteDatasetResponse struct {
	DeleteDatasetsByPK struct {
		ID *string `json:"id"`
	} `json:"delete_datasets_by_pk"`
}

// UploadDatasetImageRequest represents the payload for uploading a dataset image to S3.
type UploadDatasetImageRequest struct {
	Extension string `json:"extension"`
}

// UploadDatasetImageResponse represents the response containing presigned S3 upload details.
type UploadDatasetImageResponse struct {
	UploadDatasetImage *struct {
		Fields map[string]string `json:"fields"`
		ID     *string           `json:"id"`
		Key    *string           `json:"key"`
		URL    *string           `json:"url"`
	} `json:"uploadDatasetImage"`
}

// UploadGeneratedImageRequest represents the payload for uploading a generated image to a dataset.
type UploadGeneratedImageRequest struct {
	GeneratedImageID string `json:"generatedImageId"`
}

// UploadGeneratedImageResponse represents the response after uploading a generated image to a dataset.
type UploadGeneratedImageResponse struct {
	UploadDatasetImageFromGen *struct {
		ID *string `json:"id"`
	} `json:"uploadDatasetImageFromGen"`
}

// Image Generation-related types
type CreateGenerationRequest struct {
	Alchemy               *bool              `json:"alchemy,omitempty"`       // default true
	ContrastRatio         *float64           `json:"contrastRatio,omitempty"` // 0.1-1.0 inclusive
	ExpandedDomain        *bool              `json:"expandedDomain,omitempty"`
	FantasyAvatar         *bool              `json:"fantasyAvatar,omitempty"`
	GuidanceScale         *int               `json:"guidance_scale,omitempty"` // 1-20, recommended 7
	Height                *int               `json:"height,omitempty"`         // default 768
	HighContrast          *bool              `json:"highContrast,omitempty"`
	HighResolution        *bool              `json:"highResolution,omitempty"`
	ImagePrompts          []string           `json:"imagePrompts,omitempty"`
	ImagePromptWeight     *float64           `json:"imagePromptWeight,omitempty"`
	InitGenerationImageID *string            `json:"init_generation_image_id,omitempty"`
	InitImageID           *string            `json:"init_image_id,omitempty"`
	InitStrength          *float64           `json:"init_strength,omitempty"`
	ModelID               *string            `json:"modelId,omitempty"` // default b24e16ff-06e3-43eb-8d33-4416c2d75876
	NegativePrompt        *string            `json:"negative_prompt,omitempty"`
	NumImages             *int               `json:"num_images,omitempty"`          // default 4
	NumInferenceSteps     *int               `json:"num_inference_steps,omitempty"` // 10-60, default 15
	PhotoReal             *bool              `json:"photoReal,omitempty"`           // requires Alchemy=true, ModelID=nil
	PhotoRealVersion      *string            `json:"photoRealVersion,omitempty"`    // v1 or v2
	PhotoRealStrength     *float64           `json:"photoRealStrength,omitempty"`   // 0.55 for low, 0.5 for medium, 0.45 for high; default 0.55
	PresetStyle           *PresetStyle       `json:"presetStyle,omitempty"`         // default DYNAMIC
	Prompt                string             `json:"prompt"`                        // required
	PromptMagic           *bool              `json:"promptMagic,omitempty"`
	PromptMagicStrength   *float64           `json:"promptMagicStrength,omitempty"` // 0.1-1.0 inclusive
	PromptMagicVersion    *string            `json:"promptMagicVersion,omitempty"`  // v2 or v3
	Public                *bool              `json:"public,omitempty"`
	Scheduler             *Scheduler         `json:"scheduler,omitempty"`  // default EULER_DISCRETE
	SDVersion             *SDVersion         `json:"sd_version,omitempty"` // default v1_5
	Seed                  *int               `json:"seed,omitempty"`
	Tiling                *bool              `json:"tiling,omitempty"`
	Transparency          *string            `json:"transparency,omitempty"` // disabled, foreground_only; default disabled
	Ultra                 *bool              `json:"ultra,omitempty"`        // requires Alchemy=false
	Unzoom                *bool              `json:"unzoom,omitempty"`       // requires UnzoomAmount and InitImageID
	UnzoomAmount          *int               `json:"unzoomAmount,omitempty"`
	UpscaleRatio          *int               `json:"upscaleRatio,omitempty"` // NOTE: ENTERPRISE ACCOUNTS ONLY
	Width                 *int               `json:"width,omitempty"`        // 32-1024; default 1024
	CanvasRequest         *bool              `json:"canvasRequest,omitempty"`
	CanvasRequestType     *CanvasRequestType `json:"canvasRequestType,omitempty"` // INPAINT, OUTPAINT, SKETCH2IMG, IMG2IMG
	CanvasInitID          *string            `json:"canvasInitId,omitempty"`
	CanvasMaskID          *string            `json:"canvasMaskId,omitempty"`
}

type PresetStyle string

const (
	PresetStyleNone             PresetStyle = "NONE"
	PresetStyleLeonardo         PresetStyle = "LEONARDO"          // requires Alchemy=false
	PresetStyleAnime            PresetStyle = "ANIME"             // requires Alchemy=true
	PresetStyleCreative         PresetStyle = "CREATIVE"          // requires Alchemy=true
	PresetStyleDynamic          PresetStyle = "DYNAMIC"           // requires Alchemy=true
	PresetStyleEnvironment      PresetStyle = "ENVIRONMENT"       // requires Alchemy=true
	PresetStyleGeneral          PresetStyle = "GENERAL"           // requires Alchemy=true
	PresetStyleIllustration     PresetStyle = "ILLUSTRATION"      // requires Alchemy=true
	PresetStylePhotography      PresetStyle = "PHOTOGRAPHY"       // requires Alchemy=true
	PresetStyleRaytraced        PresetStyle = "RAYTRACED"         // requires Alchemy=true
	PresetStyleRender3D         PresetStyle = "RENDER_3D"         // requires Alchemy=true
	PresetStyleSketchBW         PresetStyle = "SKETCH_BW"         // requires Alchemy=true
	PresetStyleSketchColor      PresetStyle = "SKETCH_COLOR"      // requires Alchemy=true
	PresetStyleStockPhoto       PresetStyle = "STOCK_PHOTO"       // requires PhotoReal=true
	PresetStyleVibrant          PresetStyle = "VIBRANT"           // requires PhotoReal=true
	PresetStyleUnprocessed      PresetStyle = "UNPROCESSED"       // requires PhotoReal=true
	PresetStyleBokeh            PresetStyle = "BOKEH"             // requires PhotoReal=true
	PresetStyleCinematic        PresetStyle = "CINEMATIC"         // requires PhotoReal=true
	PresetStyleCinematicCloseup PresetStyle = "CINEMATIC_CLOSEUP" // requires PhotoReal=true
	PresetStyleFashion          PresetStyle = "FASHION"           // requires PhotoReal=true
	PresetStyleFilm             PresetStyle = "FILM"              // requires PhotoReal=true
	PresetStyleFood             PresetStyle = "FOOD"              // requires PhotoReal=true
	PresetStyleHDR              PresetStyle = "HDR"               // requires PhotoReal=true
	PresetStyleLongExposure     PresetStyle = "LONG_EXPOSURE"     // requires PhotoReal=true
	PresetStyleMacro            PresetStyle = "MACRO"             // requires PhotoReal=true
	PresetStyleMinimalistic     PresetStyle = "MINIMALISTIC"      // requires PhotoReal=true
	PresetStyleMonochrome       PresetStyle = "MONOCHROME"        // requires PhotoReal=true
	PresetStyleMoody            PresetStyle = "MOODY"             // requires PhotoReal=true
	PresetStyleNeutral          PresetStyle = "NEUTRAL"           // requires PhotoReal=true
	PresetStylePortrait         PresetStyle = "PORTRAIT"          // requires PhotoReal=true
	PresetStyleRetro            PresetStyle = "RETRO"             // requires PhotoReal=true
)

type Scheduler string

const (
	SchedulerKLMS                   Scheduler = "KLMS"
	SchedulerEulerAncestralDiscrete Scheduler = "EULER_ANCESTRAL_DISCRETE"
	SchedulerEulerDiscrete          Scheduler = "EULER_DISCRETE"
	SchedulerDDIM                   Scheduler = "DDIM"
	SchedulerDPMSolver              Scheduler = "DPM_SOLVER"
	SchedulerPNDM                   Scheduler = "PNDM"
	SchedulerLeonardo               Scheduler = "LEONARDO"
)

type SDVersion string

const (
	SDVersionV1_5           SDVersion = "v1_5"
	SDVersionV2             SDVersion = "v2"
	SDVersionV3             SDVersion = "v3"
	SDVersionSDXL_0_8       SDVersion = "SDXL_0_8"
	SDVersionSDXL_0_9       SDVersion = "SDXL_0_9"
	SDVersionSDXL_1_0       SDVersion = "SDXL_1_0"
	SDVersionSDXL_LIGHTNING SDVersion = "SDXL_LIGHTNING"
)

type CanvasRequestType string

const (
	CanvasRequestTypeInpaint    CanvasRequestType = "INPAINT"
	CanvasRequestTypeOutpaint   CanvasRequestType = "OUTPAINT"
	CanvasRequestTypeSketch2Img CanvasRequestType = "SKETCH2IMG"
	CanvasRequestTypeImg2Img    CanvasRequestType = "IMG2IMG"
)

type CreateGenerationResponse struct {
	SDGenerationJob struct {
		APICreditCost *int    `json:"apiCreditCost"`
		GenerationID  *string `json:"generationId"`
	} `json:"sdGenerationJob"`
}

type GetGenerationResponse struct {
	GenerationsByPK struct {
		ID                  *string           `json:"id"`
		Status              *GenerationStatus `json:"status"`
		CreatedAt           *Time             `json:"createdAt"`
		Height              *int              `json:"imageHeight,omitempty"`
		ModelID             *string           `json:"modelId,omitempty"`
		NegativePrompt      *string           `json:"negativePrompt,omitempty"`
		NumInferenceSteps   *int              `json:"inferenceSteps,omitempty"`
		PhotoReal           *bool             `json:"photoReal,omitempty"`
		PhotoRealStrength   *float64          `json:"photoRealStrength,omitempty"`
		PresetStyle         *PresetStyle      `json:"presetStyle,omitempty"`
		Prompt              string            `json:"prompt"`
		PromptMagic         *bool             `json:"promptMagic,omitempty"`
		PromptMagicStrength *float64          `json:"promptMagicStrength,omitempty"`
		PromptMagicVersion  *string           `json:"promptMagicVersion,omitempty"`
		Public              *bool             `json:"public,omitempty"`
		Scheduler           *Scheduler        `json:"scheduler,omitempty"`
		SDVersion           *SDVersion        `json:"sdVersion,omitempty"`
		Seed                *int              `json:"seed,omitempty"`
		Ultra               *bool             `json:"ultra,omitempty"`
		Width               *int              `json:"imageWidth,omitempty"`
		GenerationElements  []struct {
			ID            *string `json:"id"`
			Lora          *Lora   `json:"lora"`
			WeightApplied *int    `json:"weightApplied"`
		} `json:"generation_elements"`
		GeneratedImages []struct {
			ID                              *string `json:"id"`
			GeneratedImageVariationGenerics []struct {
				ID            *string           `json:"id"`
				Status        *GenerationStatus `json:"status"`
				TransformType *TransformType    `json:"transformType"`
				URL           *string           `json:"url"`
			} `json:"generated_image_variation_generics"`
			FantasyAvatar  *bool   `json:"fantasyAvatar,omitempty"`
			ImageToVideo   *bool   `json:"imageToVideo,omitempty"`
			LikeCount      *int    `json:"likeCount,omitempty"`
			Motion         *bool   `json:"motion,omitempty"`
			MotionModel    *string `json:"motionModel,omitempty"`
			MotionMP4URL   *string `json:"motionMP4Url,omitempty"`
			MotionStrength *int    `json:"motionStrength,omitempty"`
			NSFW           *bool   `json:"nsfw,omitempty"`
			URL            *string `json:"url"`
		} `json:"generated_images"`
	} `json:"generations_by_pk"`
}

type GenerationStatus string

const (
	GenerationStatusComplete GenerationStatus = "COMPLETE"
	GenerationStatusFailed   GenerationStatus = "FAILED"
	GenerationStatusPending  GenerationStatus = "PENDING"
)

type TransformType string

const (
	TransformTypeOutpaint     TransformType = "OUTPAINT"
	TransformTypeInpaint      TransformType = "INPAINT"
	TransformTypeUpscale      TransformType = "UPSCALE"
	TransformTypeUnzoom       TransformType = "UNZOOM"
	TransformTypeNoBackground TransformType = "NOBG"
)

type Generation struct {
	ID     *string `json:"id"`
	Status *string `json:"status"`
	// Add other relevant fields as necessary
	CreatedAt *Time `json:"createdAt,omitempty"`
	// Depending on the API response, include other fields like prompt, etc.
}

type GetGenerationsByUserResponse struct {
	Generations []Generation `json:"generations"`
}

type DeleteGenerationResponse struct {
	DeleteGenerationsByPK struct {
		ID *string `json:"id"`
	} `json:"delete_generations_by_pk"`
}

// Elements-related types

// ListModelsResponse represents the response when listing models.
type ListElementsResponse struct {
	Loras []Lora `json:"loras"`
}

type Lora struct {
	AKUUID        *string `json:"akUUID"`
	BaseModel     *string `json:"baseModel"`
	CreatorName   *string `json:"creatorName"`
	Description   *string `json:"description"`
	Name          *string `json:"name"`
	URLImage      *string `json:"urlImage"`
	WeightDefault *int    `json:"weightDefault"`
	WeightMax     *int    `json:"weightMax"`
	WeightMin     *int    `json:"weightMin"`
}

// User-related types
type GetUserInfoResponse struct {
	UserDetails []struct {
		APIPlanTokenRenewalDate *string `json:"apiPlanTokenRenewalDate"`
		APIConcurrencySlots     *int    `json:"apiConcurrencySlots"`
		APIPaidTokens           *int    `json:"apiPaidTokens"`
		APISubscriptionTokens   *int    `json:"apiSubscriptionTokens"`
		PaidTokens              *int    `json:"paidTokens"`
		SubscriptionGPTTokens   *int    `json:"subscriptionGptTokens"`
		SubscriptionModelTokens *int    `json:"subscriptionModelTokens"`
		SubscriptionTokens      *int    `json:"subscriptionTokens"`
		TokenRenewalDate        *string `json:"tokenRenewalDate"`
		User                    User    `json:"user"`
	} `json:"user_details"`
}

// User represents the user information.
type User struct {
	// Define user fields as per API response
	ID       *string `json:"id"`
	Username *string `json:"username"`
	// Add other fields as necessary
}

// Models-related types

// TrainCustomModelRequest represents the payload for training a new custom model.
type TrainCustomModelRequest struct {
	DatasetID      string  `json:"datasetId"`
	Description    *string `json:"description,omitempty"`
	InstancePrompt string  `json:"instance_prompt"`
	ModelType      *string `json:"modelType,omitempty"`
	Name           string  `json:"name"`
	NSFW           *bool   `json:"nsfw,omitempty"`
	Resolution     *int    `json:"resolution,omitempty"`
	SDVersion      *string `json:"sd_Version,omitempty"`
	Strength       *string `json:"strength,omitempty"`
}

// TrainCustomModelResponse represents the response from training a new custom model.
type TrainCustomModelResponse struct {
	SDTrainingJob struct {
		APICreditCost *int    `json:"apiCreditCost"`
		CustomModelID *string `json:"customModelId"`
	} `json:"sdTrainingJob"`
}

// GetCustomModelResponse represents the response when retrieving a custom model by ID.
type GetCustomModelResponse struct {
	CustomModelsByPK struct {
		CreatedAt      *Time   `json:"createdAt"`
		Description    *string `json:"description"`
		ID             *string `json:"id"`
		InstancePrompt *string `json:"instancePrompt"`
		ModelHeight    *int    `json:"modelHeight"`
		ModelWidth     *int    `json:"modelWidth"`
		Name           *string `json:"name"`
		Public         *bool   `json:"public"`
		SDVersion      *string `json:"sdVersion"`
		Status         *string `json:"status"`
		Type           *string `json:"type"`
		UpdatedAt      *Time   `json:"updatedAt"`
	} `json:"custom_models_by_pk"`
}

// DeleteCustomModelResponse represents the response when deleting a custom model.
type DeleteCustomModelResponse struct {
	DeleteCustomModelsByPK struct {
		ID *string `json:"id"`
	} `json:"delete_custom_models_by_pk"`
}

// ListPlatformModelsResponse represents the response when listing platform models.
type ListPlatformModelsResponse struct {
	CustomModels []struct {
		AKUUID        *string `json:"akUUID"`
		BaseModel     *string `json:"baseModel"`
		CreatorName   *string `json:"creatorName"`
		Description   *string `json:"description"`
		ID            *string `json:"id"`
		Name          *string `json:"name"`
		URLImage      *string `json:"urlImage"`
		WeightDefault *int    `json:"weightDefault"`
		WeightMax     *int    `json:"weightMax"`
		WeightMin     *int    `json:"weightMin"`
	} `json:"custom_models"`
}

// PaginationParams defines parameters for paginated requests.
type PaginationParams struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// UpdateCustomModelRequest represents the payload for updating a custom model.
type UpdateCustomModelRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	// Add other fields as necessary
}

// UpdateCustomModelResponse represents the response after updating a custom model.
type UpdateCustomModelResponse struct {
	UpdatedCustomModelsByPk struct {
		ID          *string `json:"id"`
		Name        *string `json:"name"`
		Description *string `json:"description"`
		// Add other fields as necessary
	} `json:"updated_custom_models_by_pk"`
}

// Prompt-related types

// GenerateRandomPromptResponse represents the response from generating a random prompt.
type GenerateRandomPromptResponse struct {
	PromptGeneration *struct {
		APICreditCost *int    `json:"apiCreditCost"`
		Prompt        *string `json:"prompt"`
	} `json:"promptGeneration"`
}

// ImprovePromptRequest represents the payload for improving a prompt.
type ImprovePromptRequest struct {
	Prompt *string `json:"prompt"`
}

// ImprovePromptResponse represents the response from improving a prompt.
type ImprovePromptResponse struct {
	PromptGeneration *struct {
		APICreditCost *int    `json:"apiCreditCost"`
		Prompt        *string `json:"prompt"`
	} `json:"promptGeneration"`
}

// Pricing Calculator-related types

// CalculateAPICostRequest represents the payload for calculating API cost.
type CalculateAPICostRequest struct {
	Service       string                 `json:"service"`
	ServiceParams map[string]interface{} `json:"serviceParams,omitempty"`
}

// CalculateAPICostResponse represents the response from calculating API cost.
type CalculateAPICostResponse struct {
	CalculateProductionApiServiceCost struct {
		Cost *int `json:"cost"`
	} `json:"calculateProductionApiServiceCost"`
}

// Realtime Canvas-related types

// CreateLCMGenerationRequest represents the payload for creating a LCM generation.
type CreateLCMGenerationRequest struct {
	Guidance         *float64 `json:"guidance,omitempty"`
	Height           *int     `json:"height,omitempty"`
	ImageDataURL     string   `json:"imageDataUrl"`
	Prompt           string   `json:"prompt"`
	RefineCreative   *bool    `json:"refineCreative,omitempty"`
	RefineStrength   *float64 `json:"refineStrength,omitempty"`
	RequestTimestamp *string  `json:"requestTimestamp,omitempty"`
	Seed             *int     `json:"seed,omitempty"`
	Steps            *int     `json:"steps,omitempty"`
	Strength         *float64 `json:"strength,omitempty"`
	Style            *string  `json:"style,omitempty"`
	Width            *int     `json:"width,omitempty"`
}

// LCMGenerationJob represents a LCM generation job.
type LCMGenerationJob struct {
	APICreditCost    *int     `json:"apiCreditCost"`
	ImageDataURL     []string `json:"imageDataUrl"`
	RequestTimestamp *string  `json:"requestTimestamp"`
}

// CreateLCMGenerationResponse represents the response after creating a LCM generation.
type CreateLCMGenerationResponse struct {
	LCMGenerationJob *LCMGenerationJob `json:"lcmGenerationJob"`
}

// PerformInstantRefineRequest represents the payload for performing instant refine.
type PerformInstantRefineRequest struct {
	Guidance         *float64 `json:"guidance,omitempty"`
	Height           *int     `json:"height,omitempty"`
	ImageDataURL     string   `json:"imageDataUrl"`
	Prompt           string   `json:"prompt"`
	RefineCreative   *bool    `json:"refineCreative,omitempty"`
	RefineStrength   *float64 `json:"refineStrength,omitempty"`
	RequestTimestamp *string  `json:"requestTimestamp,omitempty"`
	Seed             *int     `json:"seed,omitempty"`
	Steps            *int     `json:"steps,omitempty"`
	Strength         *float64 `json:"strength,omitempty"`
	Style            *string  `json:"style,omitempty"`
	Width            *int     `json:"width,omitempty"`
}

// PerformInstantRefineResponse represents the response after performing instant refine.
type PerformInstantRefineResponse struct {
	LCMGenerationJob *LCMGenerationJob `json:"lcmGenerationJob"`
}

// PerformInpaintingRequest represents the payload for performing inpainting.
type PerformInpaintingRequest struct {
	Guidance         *float64 `json:"guidance,omitempty"`
	Height           *int     `json:"height,omitempty"`
	ImageDataURL     string   `json:"imageDataUrl"`
	MaskDataURL      string   `json:"maskDataUrl"`
	Prompt           string   `json:"prompt"`
	RequestTimestamp *string  `json:"requestTimestamp,omitempty"`
	Seed             *int     `json:"seed,omitempty"`
	Steps            *int     `json:"steps,omitempty"`
	Strength         *float64 `json:"strength,omitempty"`
	Style            *string  `json:"style,omitempty"`
	Width            *int     `json:"width,omitempty"`
}

// PerformInpaintingResponse represents the response after performing inpainting.
type PerformInpaintingResponse struct {
	LCMGenerationJob *LCMGenerationJob `json:"lcmGenerationJob"`
}

// PerformAlchemyUpscaleRequest represents the payload for performing Alchemy Upscale.
type PerformAlchemyUpscaleRequest struct {
	Guidance         *float64 `json:"guidance,omitempty"`
	Height           *int     `json:"height,omitempty"`
	ImageDataURL     string   `json:"imageDataUrl"`
	Prompt           string   `json:"prompt"`
	RequestTimestamp *string  `json:"requestTimestamp,omitempty"`
	Seed             *int     `json:"seed,omitempty"`
	Steps            *int     `json:"steps,omitempty"`
	Strength         *float64 `json:"strength,omitempty"`
	Style            *string  `json:"style,omitempty"`
	Width            *int     `json:"width,omitempty"`
}

// PerformAlchemyUpscaleResponse represents the response after performing Alchemy Upscale.
type PerformAlchemyUpscaleResponse struct {
	LCMGenerationJob *struct {
		APICreditCost    *int     `json:"apiCreditCost"`
		GeneratedImageID *string  `json:"generatedImageId"`
		GenerationID     []string `json:"generationId"`
		ImageDataURL     []string `json:"imageDataUrl"`
		RequestTimestamp *string  `json:"requestTimestamp"`
		VariationID      []string `json:"variationId"`
	} `json:"lcmGenerationJob"`
}

// Texture-related types

// CreateTextureGenerationRequest represents the payload for creating a texture generation.
type CreateTextureGenerationRequest struct {
	FrontRotationOffset *int    `json:"front_rotation_offset,omitempty"`
	ModelAssetID        *string `json:"modelAssetId,omitempty"`
	NegativePrompt      *string `json:"negative_prompt,omitempty"`
	Preview             *bool   `json:"preview,omitempty"`
	PreviewDirection    *string `json:"preview_direction,omitempty"`
	Prompt              *string `json:"prompt,omitempty"`
	SDVersion           *string `json:"sd_version,omitempty"`
	Seed                *int    `json:"seed,omitempty"`
}

// CreateTextureGenerationResponse represents the response after creating a texture generation.
type CreateTextureGenerationResponse struct {
	TextureGenerationJob struct {
		APICreditCost *int    `json:"apiCreditCost"`
		ID            *string `json:"id"`
	} `json:"textureGenerationJob"`
}

// ThreeD Model Assets-related types

// Upload3DModelRequest represents the payload for uploading a 3D model.
type Upload3DModelRequest struct {
	ModelExtension *string `json:"modelExtension,omitempty"`
	Name           *string `json:"name,omitempty"`
}

// Upload3DModelResponse represents the response after uploading a 3D model.
type Upload3DModelResponse struct {
	UploadModelAsset *struct {
		ModelFields *string `json:"modelFields"`
		ModelID     *string `json:"modelId"`
		ModelKey    *string `json:"modelKey"`
		ModelURL    *string `json:"modelUrl"`
	} `json:"uploadModelAsset"`
}

// Get3DModelsByUserResponse represents the response when retrieving 3D models by user ID.
type Get3DModelsByUserResponse struct {
	ModelAssets []struct {
		CreatedAt *Time   `json:"createdAt"`
		ID        *string `json:"id"`
		MeshURL   *string `json:"meshUrl"`
		Name      *string `json:"name"`
		UpdatedAt *Time   `json:"updatedAt"`
		UserID    *string `json:"userId"`
	} `json:"model_assets"`
}

// Get3DModelByIDResponse represents the response when retrieving a 3D model by ID.
type Get3DModelByIDResponse struct {
	ModelAssetsByPK struct {
		CreatedAt   *Time   `json:"createdAt"`
		Description *string `json:"description"`
		ID          *string `json:"id"`
		MeshURL     *string `json:"meshUrl"`
		Name        *string `json:"name"`
		UpdatedAt   *Time   `json:"updatedAt"`
		UserID      *string `json:"userId"`
	} `json:"model_assets_by_pk"`
}

// Delete3DModelResponse represents the response when deleting a 3D model.
type Delete3DModelResponse struct {
	DeleteModelAssetsByPK struct {
		ID *string `json:"id"`
	} `json:"delete_model_assets_by_pk"`
}

// Variation-related types

// VariationRequest represents a generic variation request.
type VariationRequest struct {
	ID          string `json:"id"`
	IsVariation *bool  `json:"isVariation,omitempty"`
}

// UniversalUpscalerRequest represents the payload for universal upscaler.
type UniversalUpscalerRequest struct {
	ImageURL    string `json:"image_url"`
	ScaleFactor int    `json:"scale_factor"`
}

// GetVariationResponse represents the response when retrieving variation details.
type GetVariationResponse struct {
	GeneratedImageVariationGeneric []struct {
		CreatedAt     *Time   `json:"createdAt"`
		ID            *string `json:"id"`
		Status        *string `json:"status"`
		TransformType *string `json:"transformType"`
		URL           *string `json:"url"`
	} `json:"generated_image_variation_generic"`
}

// UpscaleVariationRequest represents the payload for creating an upscale variation.
type UpscaleVariationRequest struct {
	ID          string `json:"id"`
	IsVariation *bool  `json:"isVariation,omitempty"`
}

// UpscaleVariationResponse represents the response from creating an upscale variation.
type UpscaleVariationResponse struct {
	SdUpscaleJob struct {
		ID            *string `json:"id"`
		APICreditCost *int    `json:"apiCreditCost"`
	} `json:"sdUpscaleJob"`
}

// VariationJob represents variation job details.
type VariationJob struct {
	ID            *string `json:"id"`
	APICreditCost *int    `json:"apiCreditCost"`
}

// CreateUnzoomVariationResponse represents the response from creating an unzoom variation.
type CreateUnzoomVariationResponse struct {
	SdUnzoomJob VariationJob `json:"sdUnzoomJob"`
}

// CreateNoBackgroundVariationResponse represents the response from creating a no background variation.
type CreateNoBackgroundVariationResponse struct {
	SdNobgJob VariationJob `json:"sdNobgJob"`
}

// UniversalUpscalerResponse represents the response from the universal upscaler.
type UniversalUpscalerResponse struct {
	UpscaledImageURL string `json:"upscaled_image_url"`
}

// InitImages-related types
// UploadInitImageRequest represents the payload for uploading an init image.
type UploadInitImageRequest struct {
	ImageFile string `json:"image_file"`
}

// GetSingleInitImageResponse represents the response from retrieving a single init image.
type GetSingleInitImageResponse struct {
	InitImagesByPk struct {
		CreatedAt *Time   `json:"createdAt"`
		ID        *string `json:"id"`
		URL       *string `json:"url"`
	} `json:"init_images_by_pk"`
}

// DeleteInitImageResponse represents the response from deleting an init image.
type DeleteInitImageResponse struct {
	DeleteInitImagesByPk struct {
		ID *string `json:"id"`
	} `json:"delete_init_images_by_pk"`
}

// UploadCanvasInitAndMaskImageRequest represents the payload for uploading canvas init and mask images.
type UploadCanvasInitAndMaskImageRequest struct {
	InitExtension string `json:"initExtension"`
	MaskExtension string `json:"maskExtension"`
}

// UploadCanvasInitAndMaskImageResponse represents the response from uploading canvas init and mask images.
type UploadCanvasInitAndMaskImageResponse struct {
	UploadCanvasInitImage *struct {
		InitFields  string `json:"initFields"`
		InitImageID string `json:"initImageId"`
		InitKey     string `json:"initKey"`
		InitURL     string `json:"initUrl"`
		MaskFields  string `json:"maskFields"`
		MaskImageID string `json:"maskImageId"`
		MaskKey     string `json:"maskKey"`
		MaskURL     string `json:"maskUrl"`
	} `json:"uploadCanvasInitImage"`
}

// Motion-related types
// CreateSVDMotionGenerationResponse represents the response from creating an SVD motion generation.
type CreateSVDMotionGenerationResponse struct {
	Details      map[string]interface{} `json:"details"`
	GenerationID string                 `json:"generationId"`
	Status       string                 `json:"status"`
}

// CreateSVDMotionGenerationErrorResponse represents the error response from creating an SVD motion generation.
type CreateSVDMotionGenerationErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// MotionRequest represents the request payload for creating an SVD motion generation.
// Since the API specification shows empty bodyParameters, it's an empty struct.
type MotionRequest struct {
	// No fields as per API specification
}
